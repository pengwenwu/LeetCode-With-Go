package Stack

import (
	"testing"
)

func TestLinkedListPush(t *testing.T) {
	s := NewLinkedListSatck()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Print()
	s.Pop()
	s.Print()
	t.Log(s.Top())
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
}
