package queue

import "fmt"

type ArrayQueue struct {
	head int
	tail int
	capacity int
	queue []interface{}
}

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{
		head:0,
		tail:0,
		capacity:n,
		queue:make([]interface{},n)}
}

func (this *ArrayQueue) EnQueue(v interface{}) bool {
	if this.tail == this.capacity {
		return false
	}
	this.queue[this.tail] = v
	this.tail++
	return true
}

func (this *ArrayQueue) DeQueue() interface{} {
	if this.head == this.tail {
		return nil
	}
	v := this.queue[this.head]
	this.head++
	return v
}

func (this *ArrayQueue) Print() {
	if this.head == this.tail {
		fmt.Println("empty queue")
		return
	}
	result := "head"
	for i := this.head; i <= this.tail-1; i++ {
		result += fmt.Sprintf("<-%+v", this.queue[i])
	}
	result += "<-tail"
	fmt.Println(result)
}
