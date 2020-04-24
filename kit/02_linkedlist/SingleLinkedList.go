package LinkedList

import (
	"errors"
	"fmt"
)

/*
- 基本(单链表)
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
	删除链表倒数第 n 个结点
	求链表的中间结点
	在一个有环链表中，如何找出链表的入环点？以及环长
	判断两个单向链表是否相交，如果相交，求出交点
*/

// 单链表节点
type ListNode struct {
	value interface{}
	next  *ListNode
}

// 单链表
type LinkedList struct {
	head   *ListNode
	length uint
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{v, nil}
}

// 带头节点链表（哨兵）
func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

func (l *ListNode) GetNext() *ListNode {
	return l.next
}

func (l *ListNode) GetValue() interface{} {
	if l == nil {
		return errors.New("empty node")
	}
	return l.value
}

// 在某个节点之前插入新的节点
func (l *LinkedList) InsertBeforeNode(oldNode, newNode *ListNode) bool {
	if oldNode == nil || oldNode == l.head {
		return false
	}
	pre := l.head
	cur := l.head.next
	for cur != nil {
		if cur == oldNode {
			break
		}
		pre = cur
		cur = cur.next
	}
	// 未找到该节点
	if cur == nil {
		return false
	}
	newNode.next = oldNode
	pre.next = newNode
	l.length++
	return true
}

// 在某个节点之后插入新的节点
func (l *LinkedList) InsertAfterNode(oldNode, newNode *ListNode) bool {
	if oldNode == nil {
		return false
	}
	newNode.next = oldNode.next
	oldNode.next = newNode
	l.length++
	return true
}

// 头插法
func (l *LinkedList) InsertToHead(newNode *ListNode) bool {
	return l.InsertAfterNode(l.head, newNode)
}

// 尾插法
func (l *LinkedList) InsertToTail(newNode *ListNode) bool {
	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}
	return l.InsertAfterNode(cur, newNode)
}

// 通过索引查找节点
func (l *LinkedList) FindNodeByIndex(index uint) *ListNode {
	if index >= l.length {
		return nil
	}
	cur := l.head.next
	for i := uint(0); i < index; i++ {
		cur = cur.next
	}
	return cur
}

// 删除传入的节点
func (l *LinkedList) DeleteNode(node *ListNode) bool {
	pre := l.head
	cur := l.head.next
	for cur != nil {
		if cur == node {
			break
		}
		pre = cur
		cur = cur.next
	}
	// 未查询到节点
	if cur == nil {
		return false
	}
	pre.next = cur.next
	node = nil
	l.length--
	return true
}

// 打印链表
func (l *LinkedList) Print() {
	cur := l.head.next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}


// -----------------练习------------------
// 翻转链表 (非递归方法)
func (l *LinkedList) ReverseLinkedList() {
	if l.head == nil || l.head.next == nil {
		return
	}
	pre := l.head.next
	cur := l.head.next.next
	for cur != nil {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}
	// 原来第一个节点要断开
	l.head.next.next = nil
	l.head.next = pre
}
// 翻转链表（递归方法）
func (l *LinkedList) ReverseLinkedList2() *ListNode {
}
//
//// 判断是否是回文字符串
//// 思路一：找到中间点，翻转前半部分链表，时间复杂度O(n)
//func (this *LinkedList) IsPalindrome() bool {
//	switch this.length {
//	case 0:
//		return false
//	case 1:
//		return true
//	}
//	isPalindrome := true
//	step := this.length / 2
//	// 反转前半部分链表
//	var pre *ListNode = nil
//	cur := this.head.next
//	for i := uint(1); i <= step; i++ {
//		tmp := cur.next
//		cur.next = pre
//		pre = cur
//		cur = tmp
//	}
//	mid := cur
//
//	var left, right *ListNode = pre, nil
//	if this.length%2 != 0 {
//		right = mid.next
//	} else {
//		right = mid
//	}
//	for nil != left && nil != right {
//		if left.GetValue().(string) != right.GetValue().(string) {
//			isPalindrome = false
//			break
//		}
//		left = left.next
//		right = right.next
//	}
//
//	// 还原链表
//	cur = pre
//	pre = mid
//	for nil != cur {
//		tmp := cur.next
//		cur.next = pre
//		pre = cur
//		cur = tmp
//	}
//	this.head.next = pre
//	return isPalindrome
//}
//
//// 思路二：申请一个栈存储链表前半段，时间复杂度O(n)，空间复杂度O(n)
//func (this *LinkedList) isPalindrome2() bool {
//	len := this.length
//	switch len {
//	case 0:
//		return false
//	case 1:
//		return true
//	}
//	str := make([]string, 0, len/2)
//	cur := this.head
//	for i := uint(1); i <= len; i++ {
//		cur = cur.next
//		if len%2 != 0 && i == (len/2+1) { // 如果链表有奇数个节点，忽略中间节点
//			continue
//		}
//		if i <= len/2 { // 前半段入栈，后半段对比
//			str = append(str, cur.value.(string))
//		} else {
//			if str[len-i] != cur.value.(string) {
//				return false
//			}
//		}
//	}
//	return true
//}
//
//// 判断链表是否有环
//// 方法一： 双重遍历，时间复杂度O(n*n)，空间复杂度O(1)
//func (this *LinkedList) hasCycle() bool {
//	cur := this.head.next
//	for count := int(1); nil != cur; count++ {
//		new_cur := this.head.next
//		for new_count := 1; new_count < count; new_count++ {
//			if new_cur == cur {
//				return true
//			}
//			new_cur = new_cur.next
//		}
//		cur = cur.next
//	}
//	return false
//}
//
//// 方法二：HashSet存储，时间复杂度O(n)，空间复杂度O(n)
//
//// 方法三：快慢指针，时间复杂度O(n)，空间复杂度O(1)
//func (this *LinkedList) hasCycle2() bool {
//	slow := this.head.next
//	fast := this.head.next
//	for nil != fast && nil != fast.next {
//		slow = slow.next
//		fast = fast.next.next
//		if fast == slow {
//			return true
//		}
//	}
//	return false
//}
//
//// 获取有环链表入环点
//func (this *LinkedList) getEntryNodeOfLoop() interface{} {
//	slow := this.head.next
//	fast := this.head.next
//	for nil != fast && nil != fast.next {
//		slow = slow.next
//		fast = fast.next.next
//		if fast == slow {
//			break
//		}
//	}
//	// 相遇后慢指针返回链表头
//	slow = this.head.next
//	for slow != fast {
//		slow = slow.next
//		fast = fast.next
//	}
//	return slow
//}
//
//// 获取有环链表环长
//func (this *LinkedList) getCycleLengthOfLoop() uint {
//	slow := this.head.next
//	fast := this.head.next
//	for nil != fast && nil != fast.next {
//		slow = slow.next
//		fast = fast.next.next
//		if slow == fast {
//			break
//		}
//	}
//	slow = slow.next
//	fast = fast.next.next
//	var length uint = 1
//	for slow != fast {
//		slow = slow.next
//		fast = fast.next.next
//		length++
//	}
//	return length
//}
//
//// 判断两个链表是否相交，以及获取第一个交点
//// 分有环和无环
//func hasIntersection(linked_list1 *LinkedList, linked_list2 *LinkedList) interface{} {
//	// 一个有环，一个无环，必然不相交
//	// 如果都是无环，判断最后一个节点是否相同
//	if !linked_list1.hasCycle2() && !linked_list2.hasCycle2() {
//		end1 := linked_list1.head.next
//		end2 := linked_list2.head.next
//		for end1.next != nil {
//			end1 = end1.next
//		}
//		for end2.next != nil {
//			end2 = end2.next
//		}
//		if end1 == end2 {
//			return true
//		}
//		return false
//	} else if linked_list1.hasCycle2() && linked_list2.hasCycle2() {
//		// 都有环，必然是同一个环，则判断A中的入环点是否在B中出现过
//		entryNode := linked_list1.getEntryNodeOfLoop()
//		cur := linked_list2.head.next
//		for cur != nil {
//			if cur == entryNode {
//				return true
//			}
//			cur = cur.next
//		}
//		return false
//	} else {
//		return false
//	}
//}
