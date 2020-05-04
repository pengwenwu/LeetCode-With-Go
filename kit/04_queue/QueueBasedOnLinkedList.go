package queue

import "fmt"

type ListNode struct {
	next *ListNode
	value interface{}
}

type LinkedListQueue struct {
	head *ListNode
	tail *ListNode
	length int
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{
		head:nil,
		tail:nil,
		length:0,
	}
}

func (this *LinkedListQueue) EnQueue(v interface{}) bool {
	node := &ListNode{next:nil, value:v}
	if this.tail == nil {
		this.head = node
		this.tail = node
	} else {
		this.tail.next = node
		this.tail = node
	}
	this.length++
	return true
}

func (this *LinkedListQueue) DeQueue() interface{} {
	if this.head == this.tail {
		return nil
	}
	v := this.head.value
	this.head = this.head.next
	this.length--
	return v
}

func (this *LinkedListQueue) Print() {
	if this.head == nil {
		fmt.Println("empty queue")
	}
	result := "head<-"
	for cur := this.head; cur != nil; cur = cur.next {
		result += fmt.Sprintf("<-%+v", cur.value)
	}
	result += "<-tail"
	fmt.Println(result)
}