package leetcode0002

/*
Iterate through the lists adding the values and carry-over as request. Once
either of the list ends process the remaning nodes

Complexity:
Time complexity: O(n), we traverse each element of the list once
Space complexity: O(n), we need extra space to store the new list
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3, l, cur *ListNode
	carry := 0
	for (l1 != nil) && (l2 != nil) {
		val := l1.Val + l2.Val + carry
		carry = 0
		if val >= 10 {
			carry = 1
			val -= 10
		}

		l1 = l1.Next
		l2 = l2.Next
		n := &ListNode{Val: val}
		if l3 == nil {
			l3 = n
		}

		if cur != nil {
			cur.Next = n
		}

		cur = n
	}

	if l1 != nil {
		l = l1
	}
	if l2 != nil {
		l = l2
	}
	for l != nil {
		val := l.Val + carry
		carry = 0
		if val >= 10 {
			carry = 1
			val -= 10
		}

		l = l.Next
		n := &ListNode{Val: val}
		if l3 == nil {
			l3 = n
		}

		if cur != nil {
			cur.Next = n
		}

		cur = n
	}

	if carry != 0 {
		n := &ListNode{Val: carry}
		if l3 == nil {
			l3 = n
		}

		if cur != nil {
			cur.Next = n
		}

		cur = n
	}

	return l3
}
