package LinkedList

import (
	"fmt"
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

func TestIsPalindrome2(t *testing.T) {
	l := NewLinkedList()
	str := "helleh"
	for _, v := range str {
		l.InsertToTail(string(v))
	}
	l.Print()
	t.Log(l.IsPalindrome())
	l.Print()
}

func TestHasCyCle(t *testing.T) {
	node1 := NewListNode(1)
	node2 := NewListNode(2)
	node3 := NewListNode(3)
	node4 := NewListNode(4)
	node5 := NewListNode(5)
	node6 := NewListNode(6)
	node7 := NewListNode(7)
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = node6
	node6.next = node7
	node7.next = node4

	linked_list := NewLinkedList()
	linked_list.head.next = node1

	t.Log(linked_list.hasCycle())
	fmt.Println()
}

func TestHasCyCle2(t *testing.T) {
	node1 := NewListNode(1)
	node2 := NewListNode(2)
	node3 := NewListNode(3)
	node4 := NewListNode(4)
	node5 := NewListNode(5)
	node6 := NewListNode(6)
	node7 := NewListNode(7)
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = node6
	node6.next = node7
	node7.next = node4

	linked_list := NewLinkedList()
	linked_list.head.next = node1

	t.Log(linked_list.hasCycle2())
}

func TestGetEntryNodeOfLoop(t *testing.T) {
	node1 := NewListNode(1)
	node2 := NewListNode(2)
	node3 := NewListNode(3)
	node4 := NewListNode(4)
	node5 := NewListNode(5)
	node6 := NewListNode(6)
	node7 := NewListNode(7)
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = node6
	node6.next = node7
	node7.next = node4

	linked_list := NewLinkedList()
	linked_list.head.next = node1
	t.Log(linked_list.getEntryNodeOfLoop())
}

func TestGetCycleLengthOfLoop(t *testing.T) {
	node1 := NewListNode(1)
	node2 := NewListNode(2)
	node3 := NewListNode(3)
	node4 := NewListNode(4)
	node5 := NewListNode(5)
	node6 := NewListNode(6)
	node7 := NewListNode(7)
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = node6
	node6.next = node7
	node7.next = node4

	linked_list := NewLinkedList()
	linked_list.head.next = node1
	t.Log(linked_list.getCycleLengthOfLoop())
}
