package leetcode0206

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Simple one pass solution. Delete node at the head and add it to tail of the
new list. Time complexity is O(n), since we traverse the list once. Space
complexity is O(1) as no additional space is required.

Complexity:
Time complexity: O(n), traverse the list once
Space complexity: O(1), no additional space required
*/
func reverseListV0(head *ListNode) *ListNode {
	cur := head
	head = nil
	for cur != nil {
		tmp := cur
		cur = cur.Next
		tmp.Next = head
		head = tmp
	}

	return head
}

/*
Simple recursive solution. Reverse the remainder of the list and add the
current node to the tail. Keep calling till entire list is reversed. Time
complexity is O(n), since we traverse the list once. Space complexity is O(1)
as no additional space is required.

Complexity:
Time complexity: O(n), traverse the list once
Space complexity: O(1), no additional space required
*/
func reverseListV1(head *ListNode) *ListNode {
	var reverse func(*ListNode) *ListNode
	reverse = func(cur *ListNode) *ListNode {
		if cur == nil {
			return nil
		}

		node := reverse(cur.Next)
		if node == nil {
			return cur
		}

		cur.Next.Next = cur

		return node
	}

	if head == nil {
		return nil
	}

	node := reverse(head)
	head.Next = nil

	return node
}
