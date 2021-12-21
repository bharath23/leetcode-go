package leetcode0203

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name string
	head []int
	val  int
	want []int
}{
	{
		name: "test 1",
		head: []int{1, 2, 6, 3, 4, 5, 6},
		val:  6,
		want: []int{1, 2, 3, 4, 5},
	},
	{
		name: "test 2",
		head: []int{},
		val:  1,
		want: []int{},
	},
	{
		name: "test 1",
		head: []int{7, 7, 7, 7},
		val:  7,
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

func TestSolution(t *testing.T) {
	for _, test := range tests {
		head := makeList(test.head)
		have := removeElements(head, test.val)
		haveSlice := listToSlice(have)
		assert.Equalf(t, test.want, haveSlice, "%s: post deletion list mismatch", test.name)
	}
}
