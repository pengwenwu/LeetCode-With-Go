package LinkedList

import (
	"fmt"
	"testing"
)

func TestLinkedList_InsertToHead(t *testing.T) {
	l := NewLinkedList()
	total := 10
	fmt.Printf("添加一个长度为%d的链表\n", total)
	for i:= 0; i < total; i++ {
		l.InsertToTail(NewListNode(i * i))
	}
	l.Print()
	newValue := 10
	fmt.Printf("头插法插入%d\n", newValue)
	l.InsertToHead(NewListNode(newValue))
	l.Print()
}

func TestLinkedList_InsertToTail(t *testing.T) {
	l := NewLinkedList()
	total := 10
	fmt.Printf("添加一个长度为%d的链表\n", total)
	for i:= 0; i < total; i++ {
		l.InsertToTail(NewListNode(i * i))
	}
	l.Print()
	newValue := 10
	fmt.Printf("头插法插入%d\n", newValue)
	l.InsertToTail(NewListNode(newValue))
	l.Print()
}

func TestLinkedList_FindNodeByIndex(t *testing.T) {
	l := NewLinkedList()
	total := 10
	fmt.Printf("添加一个长度为%d的链表\n", total)
	for i:= 0; i < total; i++ {
		l.InsertToTail(NewListNode(i * i))
	}
	l.Print()

	newIndex := uint(9)
	fmt.Printf("查找索引为%d的节点\n", newIndex)
	fmt.Println(l.FindNodeByIndex(newIndex).GetValue())
}

func TestLinkedList_InsertAfterNode(t *testing.T) {
	l := NewLinkedList()
	total := 10
	fmt.Printf("添加一个长度为%d的链表\n", total)
	for i:= 0; i < total; i++ {
		l.InsertToTail(NewListNode(i * i))
	}
	l.Print()

	newIndex := uint(5)
	fmt.Printf("查找索引为%d的节点\n", newIndex)
	oldNode := l.FindNodeByIndex(newIndex)
	fmt.Println(oldNode.GetValue())

	newValue := 666
	fmt.Printf("后面添加值为%d的新节点\n", newValue)
	l.InsertAfterNode(oldNode, NewListNode(newValue))
	l.Print()
}

func TestLinkedList_InsertBeforeNode(t *testing.T) {
	l := NewLinkedList()
	total := 10
	fmt.Printf("添加一个长度为%d的链表\n", total)
	for i:= 0; i < total; i++ {
		l.InsertToTail(NewListNode(i * i))
	}
	l.Print()

	newIndex := uint(5)
	fmt.Printf("查找索引为%d的节点\n", newIndex)
	oldNode := l.FindNodeByIndex(newIndex)
	fmt.Println(oldNode.GetValue())

	newValue := 666
	fmt.Printf("前面添加值为%d的新节点\n", newValue)
	l.InsertBeforeNode(oldNode, NewListNode(newValue))
	l.Print()
}

func TestLinkedList_DeleteNode(t *testing.T) {
	l := NewLinkedList()
	total := 10
	fmt.Printf("添加一个长度为%d的链表\n", total)
	for i:= 0; i < total; i++ {
		l.InsertToTail(NewListNode(i * i))
	}
	l.Print()

	newIndex := uint(5)
	fmt.Printf("查找索引为%d的节点\n", newIndex)
	oldNode := l.FindNodeByIndex(newIndex)
	fmt.Println(oldNode.GetValue())

	fmt.Println("删除该节点")
	l.DeleteNode(oldNode)
	l.Print()
}

func TestLinkedList_ReverseLinkedList (t *testing.T) {
	l := NewLinkedList()
	str := "hello"
	for _, v := range str {
		l.InsertToTail(NewListNode(string(v)))
	}
	l.Print()
	l.ReverseLinkedList()
	l.Print()
}

func TestLinkedList_ReverseLinkedList2(t *testing.T) {
	l := NewLinkedList()
	str := "hello"
	for _, v := range str {
		l.InsertToTail(NewListNode(string(v)))
	}
	l.Print()
	l.ReverseLinkedList2(l.head.next)
	l.Print()
}

func TestIsPalindrome(t *testing.T) {
	l := NewLinkedList()
	str := "helloworld"
	for _, v := range str {
		l.InsertToTail(NewListNode(string(v)))
	}
	l.Print()
	fmt.Printf("%v 是否是回文字符串：%v\n", str, l.IsPalindrome())
	l.Print()
}
func TestIsPalindrome2(t *testing.T) {
	l := NewLinkedList()
	str := "helleh"
	for _, v := range str {
		l.InsertToTail(NewListNode(string(v)))
	}
	l.Print()
	fmt.Printf("%v 是否是回文字符串：%v\n", str, l.IsPalindrome())
	l.Print()
}

func TestIsPalindrome3(t *testing.T) {
	l := NewLinkedList()
	str := "h"
	for _, v := range str {
		l.InsertToTail(NewListNode(string(v)))
	}
	l.Print()
	fmt.Printf("%v 是否是回文字符串：%v\n", str, l.IsPalindrome())
	l.Print()
}

func TestLinkedList_IsPalindrome2(t *testing.T) {
	l := NewLinkedList()
	str := "helloworld"
	for _, v := range str {
		l.InsertToTail(NewListNode(string(v)))
	}
	l.Print()
	fmt.Printf("%v 是否是回文字符串：%v\n", str, l.IsPalindrome2())
	l.Print()
}

func TestLinkedList_IsPalindrome22(t *testing.T) {
	l := NewLinkedList()
	str := "heleh"
	for _, v := range str {
		l.InsertToTail(NewListNode(string(v)))
	}
	l.Print()
	fmt.Printf("%v 是否是回文字符串：%v\n", str, l.IsPalindrome2())
	l.Print()
}

func TestLinkedList_IsPalindrome23(t *testing.T) {
	l := NewLinkedList()
	str := "h"
	for _, v := range str {
		l.InsertToTail(NewListNode(string(v)))
	}
	l.Print()
	fmt.Printf("%v 是否是回文字符串：%v\n", str, l.IsPalindrome2())
	l.Print()
}

func TestLinkedList_IsPalindrome3(t *testing.T) {
	l := NewLinkedList()
	str := "heleh"
	for _, v := range str {
		l.InsertToTail(NewListNode(string(v)))
	}
	l.Print()
	fmt.Printf("%v 是否是回文字符串：%v\n", str, l.IsPalindrome3())
	l.Print()
}

func TestLinkedList_IsPalindrome32(t *testing.T) {
	l := NewLinkedList()
	str := "helloworld"
	for _, v := range str {
		l.InsertToTail(NewListNode(string(v)))
	}
	l.Print()
	fmt.Printf("%v 是否是回文字符串：%v\n", str, l.IsPalindrome3())
	l.Print()
}

func TestLinkedList_IsPalindrome33(t *testing.T) {
	l := NewLinkedList()
	str := "h"
	for _, v := range str {
		l.InsertToTail(NewListNode(string(v)))
	}
	l.Print()
	fmt.Printf("%v 是否是回文字符串：%v\n", str, l.IsPalindrome3())
	l.Print()
}

func TestLinkedList_HasCycle(t *testing.T) {
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

	l := NewLinkedList()
	l.head.next = node1
	l.Print()
	fmt.Printf("是否有环：%v\n", l.HasCycle())

	node7.next = node4
	fmt.Printf("添加一个有环节点后是否有环：%v\n", l.HasCycle())
}

func TestLinkedList_HasCycle2(t *testing.T) {
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

	l := NewLinkedList()
	l.head.next = node1
	l.Print()
	fmt.Printf("是否有环：%v\n", l.HasCycle2())

	node7.next = node4
	fmt.Printf("添加一个有环节点后是否有环：%v\n", l.HasCycle2())
}

//func TestHasCyCle2(t *testing.T) {
//	node1 := NewListNode(1)
//	node2 := NewListNode(2)
//	node3 := NewListNode(3)
//	node4 := NewListNode(4)
//	node5 := NewListNode(5)
//	node6 := NewListNode(6)
//	node7 := NewListNode(7)
//	node1.next = node2
//	node2.next = node3
//	node3.next = node4
//	node4.next = node5
//	node5.next = node6
//	node6.next = node7
//	node7.next = node4
//
//	linked_list := NewLinkedList()
//	linked_list.head.next = node1
//
//	t.Log(linked_list.hasCycle2())
//}
//
//func TestGetEntryNodeOfLoop(t *testing.T) {
//	node1 := NewListNode(1)
//	node2 := NewListNode(2)
//	node3 := NewListNode(3)
//	node4 := NewListNode(4)
//	node5 := NewListNode(5)
//	node6 := NewListNode(6)
//	node7 := NewListNode(7)
//	node1.next = node2
//	node2.next = node3
//	node3.next = node4
//	node4.next = node5
//	node5.next = node6
//	node6.next = node7
//	node7.next = node4
//
//	linked_list := NewLinkedList()
//	linked_list.head.next = node1
//	t.Log(linked_list.getEntryNodeOfLoop())
//}
//
//func TestGetCycleLengthOfLoop(t *testing.T) {
//	node1 := NewListNode(1)
//	node2 := NewListNode(2)
//	node3 := NewListNode(3)
//	node4 := NewListNode(4)
//	node5 := NewListNode(5)
//	node6 := NewListNode(6)
//	node7 := NewListNode(7)
//	node1.next = node2
//	node2.next = node3
//	node3.next = node4
//	node4.next = node5
//	node5.next = node6
//	node6.next = node7
//	node7.next = node4
//
//	linked_list := NewLinkedList()
//	linked_list.head.next = node1
//	t.Log(linked_list.getCycleLengthOfLoop())
//}
//
//func TestHasIntersection(t *testing.T) {
//	node1 := NewListNode(1)
//	node2 := NewListNode(2)
//	node3 := NewListNode(3)
//	node4 := NewListNode(4)
//	node5 := NewListNode(5)
//	node6 := NewListNode(6)
//	node7 := NewListNode(7)
//	node1.next = node2
//	node2.next = node3
//	node3.next = node4
//	node4.next = node5
//	node5.next = node6
//	node6.next = node7
//	// node7.next = node4 // 有环1->2->3->4->5->6->7->4
//
//	linked_list := NewLinkedList()
//	linked_list.head.next = node1
//
//	node8 := NewListNode(8) // 8-9->10->6
//	node9 := NewListNode(9)
//	node10 := NewListNode(10)
//	node8.next = node9
//	node9.next = node10
//	node10.next = node6
//	linked_list2 := NewLinkedList()
//	linked_list2.head.next = node8
//
//	// linked_list.Print()
//	// linked_list2.Print()
//	t.Log(hasIntersection(linked_list, linked_list2))
//}
