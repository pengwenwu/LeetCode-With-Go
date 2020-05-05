package queue

import (
	"fmt"
	"testing"
)

func TestCircularQueue_EnQueue(t *testing.T) {
	q := NewCircularQueue(5)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	q.Print()
}

func TestCircularQueue_Dequeue(t *testing.T) {
	q := NewCircularQueue(5)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)

	q.Print()
	fmt.Println(q.DeQueue())
	q.Print()
	q.EnQueue(5)
	q.Print()
	q.DeQueue()
	q.Print()
	q.DeQueue()
	q.Print()
	q.DeQueue()
	q.Print()
	fmt.Println(q.DeQueue())
	q.Print()
}