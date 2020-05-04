package queue

import "testing"

func TestLinkedListQueue_EnQueue(t *testing.T) {
	l := NewLinkedListQueue()
	l.EnQueue("h")
	l.EnQueue("e")
	l.EnQueue("l")
	l.EnQueue("l")
	l.EnQueue("o")

	l.Print()
}

func TestLinkedListQueue_DeQueue(t *testing.T) {
	l := NewLinkedListQueue()
	l.EnQueue("h")
	l.EnQueue("e")
	l.EnQueue("l")
	l.EnQueue("l")
	l.EnQueue("o")

	l.Print()

	l.DeQueue()
	l.Print()
}