package queue

import (
	"fmt"
	"testing"
)

func TestArrayQueue_EnQueue(t *testing.T) {
	q := NewArrayQueue(10)
	q.EnQueue("h")
	q.EnQueue("e")
	q.EnQueue("l")
	q.EnQueue("l")
	q.EnQueue("o")
	q.Print()
}

func TestArrayQueue_DeQueue(t *testing.T) {
	q := NewArrayQueue(10)
	q.EnQueue("h")
	q.EnQueue("e")
	q.EnQueue("l")
	q.EnQueue("l")
	q.EnQueue("o")

	q.Print()
	fmt.Println(q.DeQueue())
	q.Print()
}