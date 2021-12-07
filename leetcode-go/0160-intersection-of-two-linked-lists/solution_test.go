package leetcode0160

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name         string
	intersectVal int
	listA        []int
	listB        []int
	skipA        int
	skipB        int
}{
	{
		name:         "test 1",
		intersectVal: 8,
		listA:        []int{4, 1, 8, 4, 5},
		listB:        []int{5, 6, 1, 8, 4, 5},
		skipA:        2,
		skipB:        3,
	},
	{
		name:         "test 2",
		intersectVal: 2,
		listA:        []int{1, 9, 1, 2, 4},
		listB:        []int{3, 2, 4},
		skipA:        3,
		skipB:        1,
	},
	{
		name:         "test 3",
		intersectVal: 0,
		listA:        []int{2, 6, 4},
		listB:        []int{1, 5},
		skipA:        3,
		skipB:        2,
	},
}

func makeLists(intersectVal, skipA, skipB int, listA, listB []int) (*ListNode, *ListNode, *ListNode) {
	var common, headA, headB, tailA, tailB *ListNode
	for i, v := range listA {
		n := &ListNode{Val: v}
		if headA == nil {
			headA = n
		}

		if tailA != nil {
			tailA.Next = n
		}

		tailA = n
		if (i == skipA) && (intersectVal == v) {
			common = n
		}
	}

	for i, v := range listB {
		if (i == skipB) && (intersectVal == v) {
			if headB == nil {
				headB = common
			}

			if tailB != nil {
				tailB.Next = common
			}

			tailB = tailA
			break
		}

		n := &ListNode{Val: v}
		if headB == nil {
			headB = n
		}

		if tailB != nil {
			tailB.Next = n
		}

		tailB = n
	}

	return headA, headB, common
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		headA, headB, want := makeLists(test.intersectVal, test.skipA, test.skipB, test.listA, test.listB)
		have := getIntersectionNodeV0(headA, headB)
		assert.Equalf(t, want, have, "%s: list intersection does not match", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		headA, headB, want := makeLists(test.intersectVal, test.skipA, test.skipB, test.listA, test.listB)
		have := getIntersectionNodeV1(headA, headB)
		assert.Equalf(t, want, have, "%s: list intersection does not match", test.name)
	}
}

func TestSolutionV2(t *testing.T) {
	for _, test := range tests {
		headA, headB, want := makeLists(test.intersectVal, test.skipA, test.skipB, test.listA, test.listB)
		have := getIntersectionNodeV2(headA, headB)
		assert.Equalf(t, want, have, "%s: list intersection does not match", test.name)
	}
}
