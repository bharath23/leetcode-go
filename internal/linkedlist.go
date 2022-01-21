package internal

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewCircularListFromSlice(input []int) *ListNode {
	var head, tail *ListNode
	for _, v := range input {
		n := &ListNode{Val: v}
		if head == nil {
			head = n
		}

		n.Next = head
		if tail == nil {
			tail = n
		}

		tail.Next = n
		tail = n
	}

	return head
}

func NewListFromSlice(input []int) *ListNode {
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

func NewListFromSliceWithCycle(input []int, pos int) (*ListNode, *ListNode) {
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

func NewListsWithIntersection(val, skipA, skipB int, inputA, inputB []int) (*ListNode, *ListNode, *ListNode) {
	var headA, tailA, headB, tailB, intersect *ListNode
	for i, v := range inputA {
		n := &ListNode{Val: v}
		if headA == nil {
			headA = n
		}

		if tailA != nil {
			tailA.Next = n
		}

		tailA = n
		if i == skipA {
			intersect = n
		}
	}

	for i, v := range inputB {
		if (i == skipB) && (v == val) {
			if headB == nil {
				headB = intersect
			}

			if tailB != nil {
				tailB.Next = intersect
			}

			tailB = tailA
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
	return headA, headB, intersect
}

func ListToSlice(head *ListNode) []int {
	s := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		s = append(s, cur.Val)
	}

	return s
}

func CircularListToSlice(head *ListNode) []int {
	s := []int{}
	cur := head
	if cur != nil {
		s = append(s, cur.Val)
		cur = cur.Next
	}

	for cur != nil && cur != head {
		s = append(s, cur.Val)
		cur = cur.Next
	}

	return s
}
