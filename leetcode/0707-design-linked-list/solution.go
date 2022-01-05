package leetcode0707

type ListNode struct {
	val  int
	next *ListNode
	prev *ListNode
}

type MyLinkedList struct {
	head *ListNode
	tail *ListNode
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{head: nil, tail: nil, size: 0}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index > this.size-1 {
		return -1
	}

	cur := this.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}

	return cur.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	n := &ListNode{val: val}
	n.next = this.head
	if this.head != nil {
		this.head.prev = n
	}

	this.head = n
	this.size++
	if this.tail == nil {
		this.tail = n
	}
}

func (this *MyLinkedList) AddAtTail(val int) {
	n := &ListNode{val: val}
	if this.tail != nil {
		this.tail.next = n
	}

	n.prev = this.tail
	this.tail = n
	this.size++
	if this.head == nil {
		this.head = n
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	}

	if index == this.size {
		this.AddAtTail(val)
		return
	}

	n := &ListNode{val: val}
	cur := this.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}

	if cur != nil {
		n.prev = cur.prev
		cur.prev = n
	}

	if n.prev != nil {
		n.prev.next = n
	}

	n.next = cur
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > this.size-1 {
		return
	}

	cur := this.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}

	if cur.prev != nil {
		cur.prev.next = cur.next
	}
	if cur.next != nil {
		cur.next.prev = cur.prev
	}

	if this.head == cur {
		this.head = cur.next
	}

	if this.tail == cur {
		this.tail = cur.prev
	}

	this.size--
}
