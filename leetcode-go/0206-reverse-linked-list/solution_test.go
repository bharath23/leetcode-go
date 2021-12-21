package leetcode0206

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name string
	head []int
	want []int
}{
	{
		name: "test 1",
		head: []int{1, 2, 3, 4, 5},
		want: []int{5, 4, 3, 2, 1},
	},
	{
		name: "test 2",
		head: []int{1, 2},
		want: []int{2, 1},
	},
	{
		name: "test 3",
		head: []int{},
		want: []int{},
	},
}

func makeList(input []int) *ListNode {
	var head, tail *ListNode
	for _, v := range input {
		n := &ListNode{Val: v}
		if head == nil {
			head = n
		}

		if tail != nil {
			tail.Next = n
		}

		tail = n
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
		have := reverseListV0(head)
		haveSlice := listToSlice(have)
		assert.Equalf(t, test.want, haveSlice, "%s: reverse list does not match", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		head := makeList(test.head)
		have := reverseListV1(head)
		haveSlice := listToSlice(have)
		assert.Equalf(t, test.want, haveSlice, "%s: reverse list does not match", test.name)
	}
}
