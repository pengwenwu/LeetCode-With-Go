package LinkedList

import (
	"testing"
)

func TestInsertToHead(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToHead(i * i)
	}
	l.Print()
}

func TestInsertToTail(t *testing.T) {
	l := NewLinkedList()
	str := "hello"
	for _, v := range str {
		l.InsertToTail(string(v))
	}
	l.Print()
}

func TestInsertBefore(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToHead(i * i)
	}
	oldNode := NewListNode("oldNode")
	l.InsertToHead(oldNode)
	l.Print()
}

func TestInsertAfter(t *testing.T) {

}

func TestFindByIndex(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToHead(i * i)
	}
	t.Log(l.FindByIndex(2))
}

func TestDeleteNode(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 3; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()

	t.Log(l.DeleteNode(l.head.next))
	l.Print()

	t.Log(l.DeleteNode(l.head.next.next))
	l.Print()
}

func TestReverseLinkedList(t *testing.T) {
	l := NewLinkedList()
	str := "hello"
	for _, v := range str {
		l.InsertToTail(string(v))
	}
	l.Print()
	l.ReverseLinkedList()
	l.Print()
}

func TestIsPalindrome(t *testing.T) {
	l := NewLinkedList()
	str := "helloworld"
	for _, v := range str {
		l.InsertToTail(string(v))
	}
	l.Print()
	t.Log(l.IsPalindrome())
	l.Print()
}
