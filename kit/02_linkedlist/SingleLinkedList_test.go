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

func TestLinkedList_GetEntryNodeOfLoop(t *testing.T) {
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

	node7.next = node4
	fmt.Printf("添加一个有环节点后节点值：%v\n", l.GetEntryNodeOfLoop().value)
}

func TestLinkedList_GetEntryNodeOfLoop2(t *testing.T) {
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

	node7.next = node2
	fmt.Printf("添加一个有环节点后节点值：%v\n", l.GetEntryNodeOfLoop2().value)
}

func TestLinkedList_GetLengthToEntry(t *testing.T) {
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

	node7.next = node2
	fmt.Printf("添加一个有环节点后到入口柄长：%v\n", l.GetLengthToEntry())
}

func TestLinkedList_GetLengthOfLoop(t *testing.T) {
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

	node7.next = node5
	fmt.Printf("添加一个有环节点后环长：%v\n", l.GetLengthOfLoop())
}

func TestLinkedList_HasIntersection(t *testing.T) {
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
	node7.next = node4 // 有环1->2->3->4->5->6->7->4

	l1 := NewLinkedList()
	l1.head.next = node1

	node8 := NewListNode(8)
	node9 := NewListNode(9)
	node10 := NewListNode(10)
	node8.next = node9
	node9.next = node10
	node10.next = node6 // 8-9->10->6 // 入环点在环内
	//node10.next = node3 // 8-9->10->3 // 入环点在环外

	l2 := NewLinkedList()
	l2.head.next = node8

	//l1.Print()
	//l2.Print()
	fmt.Printf("l1是否有环：%+v\n", l1.HasCycle2())
	fmt.Printf("l2是否有环：%+v\n", l2.HasCycle2())

	fmt.Printf("l1, l2是否相交：%+v\n", l1.HasIntersection(l1, l2))
}

func TestLinkedList_MergeSortLinkedList(t *testing.T) {
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
	node6.next = node7  // 1-2-3-4-5-6-7

	l1 := NewLinkedList()
	l1.head.next = node1

	node8 := NewListNode(8)
	node9 := NewListNode(9)
	node10 := NewListNode(10)
	node8.next = node9
	node9.next = node10 // 8-9-10

	l2 := NewLinkedList()
	l2.head.next = node8

	l1.Print()
	l2.Print()
	l3 := l1.MergeSortLinkedList(l1, l2)
	l3.Print()
}

func TestLinkedList_MergeSortLinkedList2(t *testing.T) {
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
	node6.next = node7  // 1-2-3-4-5-6-7

	l1 := NewLinkedList()
	l1.head.next = node1

	node11 := NewListNode(-10)
	node12 := NewListNode(0)
	node13 := NewListNode(1)
	node14 := NewListNode(6)
	node15 := NewListNode(8)
	node11.next = node12
	node12.next = node13
	node13.next = node14
	node14.next = node15 // -10-0-1-6-8

	l2 := NewLinkedList()
	l2.head.next = node11

	l1.Print()
	l2.Print()
	l3 := l1.MergeSortLinkedList(l1, l2)
	l3.Print()
}
