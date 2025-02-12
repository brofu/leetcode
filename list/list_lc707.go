package list

import "fmt"

type IntNode struct {
	prev  *IntNode
	next  *IntNode
	value int
}

type MyLinkedList struct {
	head *IntNode
	tail *IntNode
	size int
}

func Constructor() MyLinkedList {

	head := &IntNode{}
	tail := &IntNode{
		prev: head,
	}
	head.next = tail
	return MyLinkedList{
		head: head,
		tail: tail,
		size: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	p := this.head
	for idx := 0; idx <= index; idx++ {
		p = p.next
	}
	return p.value
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}

	p := this.head
	for idx := 0; idx < index; idx++ {
		p = p.next
	}
	node := &IntNode{
		value: val,
	}
	node.prev = p
	node.next = p.next
	p.next.prev = node
	p.next = node

	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}

	p := this.head
	for idx := 0; idx < index; idx++ {
		p = p.next
	}
	node := p.next
	p.next = node.next
	node.next.prev = p

	this.size--
}

func (this *MyLinkedList) Debug(message string) {
	p := this.head
	fmt.Println(message)
	for p.next != nil && p.next != this.tail {
		p = p.next
		fmt.Println(p.value)
	}
}
