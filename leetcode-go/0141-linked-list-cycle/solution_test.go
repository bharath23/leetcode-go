package leetcode0141

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name  string
	input []int
	pos   int
	want  bool
}{
	{
		name:  "test 1",
		input: []int{3, 2, 0, -4},
		pos:   1,
		want:  true,
	},
	{
		name:  "test 2",
		input: []int{1, 2},
		pos:   0,
		want:  true,
	},
	{
		name:  "test 3",
		input: []int{1},
		pos:   -1,
		want:  false,
	},
	{
		name:  "test 4",
		input: []int{},
		pos:   -1,
		want:  false,
	},
}

func makeList(input []int, pos int) *ListNode {
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

	return head
}

func TestSolutionV0(t *testing.T) {
	for _, test := range tests {
		head := makeList(test.input, test.pos)
		have := hasCycleV0(head)
		assert.Equalf(t, test.want, have, "%s: testing for cycle mismatch", test.name)
	}
}

func TestSolutionV1(t *testing.T) {
	for _, test := range tests {
		head := makeList(test.input, test.pos)
		have := hasCycleV1(head)
		assert.Equalf(t, test.want, have, "%s: testing for cycle mismatch", test.name)
	}
}
