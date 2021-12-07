package leetcode0141

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name  string
	input []int
	pos   int
}{
	{
		name:  "test 1",
		input: []int{3, 2, 0, -4},
		pos:   1,
	},
	{
		name:  "test 2",
		input: []int{1, 2},
		pos:   0,
	},
	{
		name:  "test 3",
		input: []int{1},
		pos:   -1,
	},
	{
		name:  "test 4",
		input: []int{},
		pos:   -1,
	},
}

func makeList(input []int, pos int) (*ListNode, *ListNode) {
	var cycle, head, tail *ListNode
	for i, v := range input {
		n := &ListNode{Val: v}
		if head == nil {
			head = n
		}

		if tail != nil {
			tail.Next = n
		}
		tail = n
		if i == pos {
			cycle = n
		}
	}

	if tail != nil {
		tail.Next = cycle
	}

	return head, cycle
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		head, want := makeList(test.input, test.pos)
		have := detectCycleV0(head)
		assert.Equalf(t, want, have, "%s: testing for cycle mismatch", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		head, want := makeList(test.input, test.pos)
		have := detectCycleV1(head)
		assert.Equalf(t, want, have, "%s: testing for cycle mismatch", test.name)
	}
}
