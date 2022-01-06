package internal

type Node struct {
	Val   int
	Child *Node
	Next  *Node
	Prev  *Node
}

func NewMultiLevelDoublyLinkedList(input []int) *Node {
	newDoublyLinkedList := func(input []int, start, end int) (*Node, int) {
		var head, tail *Node
		i := start
		for i < end {
			if input[i] == -1 {
				i++
				break
			}
			n := &Node{Val: input[i]}
			if head == nil {
				head = n
			}

			n.Prev = tail
			if tail != nil {
				tail.Next = n
			}

			tail = n
			i++
		}

		return head, i
	}

	var head, cur *Node
	end := len(input)
	i := 0
	for i < end {
		if input[i] == -1 {
			cur = cur.Next
			i++
		} else {
			var node *Node
			node, i = newDoublyLinkedList(input, i, end)
			if head == nil {
				head = node
			}
			if cur != nil {
				cur.Child = node
			}

			cur = node
		}
	}

	return head
}

func MultiLevelDoublyLinkedListToSlice(head *Node) []int {
	s := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		s = append(s, cur.Val)
	}

	return s
}
