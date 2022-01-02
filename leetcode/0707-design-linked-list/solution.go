package leetcode0707

//import "fmt"

type ListNode struct {
	val  int
	next *ListNode
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
	n := ListNode{val: val}
	n.next = this.head
	this.head = &n
	this.size++
	if this.tail == nil {
		this.tail = &n
	}
}

func (this *MyLinkedList) AddAtTail(val int) {
	n := ListNode{val: val}
	if this.tail != nil {
		this.tail.next = &n
	}

	this.tail = &n
	this.size++
	if this.head == nil {
		this.head = &n
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}

	n := &ListNode{val: val}
	var prev *ListNode
	cur := this.head
	for i := 0; i < index; i++ {
		prev = cur
		cur = cur.next
	}

	n.next = cur
	if prev != nil {
		prev.next = n
	}

	if index == 0 {
		this.head = n
	}

	if index == this.size {
		this.tail = n
	}
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > this.size-1 {
		return
	}

	var prev *ListNode
	cur := this.head
	for i := 0; i < index; i++ {
		prev = cur
		cur = cur.next
	}

	if prev != nil {
		prev.next = cur.next
	}

	if this.head == cur {
		this.head = cur.next
	}

	if this.tail == cur {
		this.tail = prev
	}

	this.size--
}
