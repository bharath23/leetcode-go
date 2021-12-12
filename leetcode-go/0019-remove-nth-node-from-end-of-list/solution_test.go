package leetcode0019

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name string
	head []int
	n    int
	want []int
}{
	{
		name: "test 1",
		head: []int{1, 2, 3, 4, 5},
		n:    2,
		want: []int{1, 2, 3, 5},
	},
}

func makeList(input []int) *ListNode {
	var head, tail *ListNode
	for _, v := range input {
		n := &ListNode{Val: v}
		if tail != nil {
			tail.Next = n
		}

		tail = n
		if head == nil {
			head = n
		}
	}

	return head
}

func listToSlice(head *ListNode) []int {
	s := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		s = append(s, cur.Val)
	}

	return s
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		head := makeList(test.head)
		have := removeNthFromEndV0(head, test.n)
		assert.Equalf(t, test.want, listToSlice(have), "%s: list different", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		head := makeList(test.head)
		have := removeNthFromEndV1(head, test.n)
		assert.Equalf(t, test.want, listToSlice(have), "%s: list different", test.name)
	}
}
