package queue

import "fmt"

type CircularQueue struct {
	head     int
	tail     int
	capacity int
	queue    []interface{}
}

func NewCircularQueue(n int) *CircularQueue {
	return &CircularQueue{
		head:     0,
		tail:     0,
		capacity: n,
		queue:    make([]interface{}, n),
	}
}

func (this *CircularQueue) EnQueue(v interface{}) bool {
	if this.isFull() {
		return false
	}
	this.queue[this.tail] = v
	this.tail = (this.tail + 1) % this.capacity
	return true
}

func (this *CircularQueue) DeQueue() interface{} {
	if this.isEmpty() {
		return nil
	}
	v := this.queue[this.head]
	this.head = (this.head + 1) % this.capacity
	return v
}

func (this *CircularQueue) Print() {
	if this.isEmpty() {
		fmt.Println("empty queue")
		return
	}
	result := "head"
	var i = this.head
	for true {
		result += fmt.Sprintf("<-%+v", this.queue[i])
		i = (i + 1) % this.capacity
		if i == this.tail {
			break
		}
	}
	result += "<-tail"
	fmt.Println(result)
}

func (this *CircularQueue) isEmpty() bool {
	if this.head == this.tail {
		return true
	}
	return false
}

func (this *CircularQueue) isFull() bool {
	if (this.tail+1)%this.capacity == this.head {
		return true
	}
	return false
}
