package leetcode0002

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func (l *ListNode) String() string {
	if l == nil {
		return ""
	}

	var s strings.Builder
	fmt.Fprintf(&s, "%d", l.Val)
	l = l.Next
	for l != nil {
		fmt.Fprintf(&s, " -> %d", l.Val)
		l = l.Next
	}

	return s.String()
}

func equalListNode(l1, l2 *ListNode) bool {
	for (l1 != nil) && (l2 != nil) {
		if l1.Val != l2.Val {
			return false
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	return (l1 == l2)
}

func makeListNode(ints []int) *ListNode {
	if len(ints) == 0 {
		return nil
	}

	root := &ListNode{
		Val:  ints[0],
		Next: nil,
	}

	cur := root
	for _, i := range ints[1:] {
		n := &ListNode{
			Val:  i,
			Next: nil,
		}
		cur.Next = n
		cur = n
	}

	return root
}

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			name: "test 1",
			l1:   makeListNode([]int{2, 4, 3}),
			l2:   makeListNode([]int{5, 6, 4}),
			want: makeListNode([]int{7, 0, 8}),
		},
		{
			name: "test 2",
			l1:   makeListNode([]int{5}),
			l2:   makeListNode([]int{5}),
			want: makeListNode([]int{0, 1}),
		},
	}

	opt := cmp.Comparer(equalListNode)
	for _, test := range tests {
		have := addTwoNumbers(test.l1, test.l2)
		assert.Truef(
			t,
			cmp.Equal(test.want, have, opt),
			"%s: sum do not match",
		)
	}
}
