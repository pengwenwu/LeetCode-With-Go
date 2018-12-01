package LinkedList

import (
	"fmt"
)

/*
- 基本
	头插法
	尾插法
	在某个节点前插入
	在某个节点后面插入
	通过索引查找节点
	删除传入的节点
	打印链表
- 练习
	单链表反转
	回文字符串判定
	链表中环的检测
	两个有序的链表合并
*/

type ListNode struct {
	next  *ListNode
	value interface{}
}

type LinkedList struct {
	head   *ListNode
	length uint
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{next: nil, value: v}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: NewListNode(0), length: 0}
}

func (this *ListNode) GetNext() *ListNode {
	return this.next
}

func (this *ListNode) GetValue() interface{} {
	return this.value
}

// 在某个节点前插入
func (this *LinkedList) InsertBefore(p *ListNode, v *ListNode) bool {
	if nil == p || p == this.head {
		return false
	}
	pre := this.head
	cur := this.head.next
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	newNode := NewListNode(v)
	pre.next = newNode
	newNode.next = cur
	this.length++
	return true
}

// 在某个节点后插入
func (this *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if nil == p {
		return false
	}
	newNode := NewListNode(v)
	oldNext := p.next
	p.next = newNode
	newNode.next = oldNext
	this.length++
	return true
}

// 头插法
func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

// 尾插法
func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for nil != cur.next {
		cur = cur.next
	}
	return this.InsertAfter(cur, v)
}

// 根据索引查找节点
func (this *LinkedList) FindByIndex(index uint) *ListNode {
	if index > this.length {
		return nil
	}
	cur := this.head.next
	for i := uint(0); i < index; i++ {
		cur = cur.next
	}
	return cur
}

// 删除节点
func (this *LinkedList) DeleteNode(p *ListNode) bool {
	if nil == p {
		return false
	}
	pre := this.head
	cur := this.head.next
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	pre.next = p.next
	p = nil
	this.length--
	return true
}

func (this *LinkedList) Print() {
	cur := this.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}
