package Stack

import (
	"testing"
)

func TestArrayStackPush(t *testing.T) {
	s := NewArrayStack()
	s.Push(1)
	s.Push(2)
	s.Print()
	t.Log(s.Pop())
	s.Push(4)
	s.Print()
}

func TestArrayStackPop(t *testing.T) {
	s := NewArrayStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Print()

	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
	s.Print()
}
