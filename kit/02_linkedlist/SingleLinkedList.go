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

// 判断是否是回文字符串
// 思路一：找到中间点，翻转前半部分链表，时间复杂度O(n)
func (this *LinkedList) IsPalindrome() bool {
	switch this.length {
	case 0:
		return false
	case 1:
		return true
	}
	isPalindrome := true
	step := this.length / 2
	// 反转前半部分链表
	var pre *ListNode = nil
	cur := this.head.next
	for i := uint(1); i <= step; i++ {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}
	mid := cur

	var left, right *ListNode = pre, nil
	if this.length%2 != 0 {
		right = mid.next
	} else {
		right = mid
	}
	for nil != left && nil != right {
		if left.GetValue().(string) != right.GetValue().(string) {
			isPalindrome = false
			break
		}
		left = left.next
		right = right.next
	}

	// 还原链表
	cur = pre
	pre = mid
	for nil != cur {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}
	this.head.next = pre
	return isPalindrome
}

// 思路二：申请一个栈存储链表前半段，时间复杂度O(n)，空间复杂度O(n)
func (this *LinkedList) isPalindrome2() bool {
	len := this.length
	switch len {
	case 0:
		return false
	case 1:
		return true
	}
	str := make([]string, 0, len/2)
	cur := this.head
	for i := uint(1); i <= len; i++ {
		cur = cur.next
		if len%2 != 0 && i == (len/2+1) { // 如果链表有奇数个节点，忽略中间节点
			continue
		}
		if i <= len/2 { // 前半段入栈，后半段对比
			str = append(str, cur.value.(string))
		} else {
			if str[len-i] != cur.value.(string) {
				return false
			}
		}
	}
	return true
}

// 翻转链表
func (this *LinkedList) ReverseLinkedList() {
	if nil == this.head || nil == this.head.next || nil == this.head.next.next {
		return
	}
	pre := this.head.next
	cur := this.head.next.next
	pre.next = nil
	for nil != cur {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}
	this.head.next = pre
}

// 判断链表是否有环
// 方法一： 双重遍历，时间复杂度O(n*n)，空间复杂度O(1)
func (this *LinkedList) hasCycle() bool {
	if this.length <= 2 {
		return false
	}
	cur := this.head.next
	for count := int(1); nil != cur; count++ {
		new_cur := this.head.next
		for new_count := 1; new_count < count; new_count++ {
			if new_cur == cur {
				return true
			}
			new_cur = new_cur.next
		}
		cur = cur.next
	}
	return false
}
