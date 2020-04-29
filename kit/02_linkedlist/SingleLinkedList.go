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
	在一个有环链表中，如何找出链表的入环点？以及环长，柄长
	判断两个单向链表是否相交，如果相交，求出交点
	两个有序的链表合并
	删除链表倒数第 n 个结点
	求链表的中间结点
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
// 翻转链表 (迭代方法)
func (l *LinkedList) ReverseLinkedList() {
	if l.head == nil || l.head.next == nil {
		return
	}
	var pre *ListNode
	cur := l.head.next
	for cur != nil {
		cur.next, pre, cur = pre, cur, cur.next
	}
	l.head.next = pre
}

// 翻转链表（递归方法）
func (l *LinkedList) ReverseLinkedList2(node *ListNode) *ListNode {
	if node == nil || node.next == nil {
		l.head.next = node
		return node
	}
	newHead := l.ReverseLinkedList2(node.next)
	node.next.next = node
	node.next = nil
	return newHead
}

// 回文链表判断
// 方法一：借用栈存储 时间复杂度O(n), 空间复杂度O(n)
func (l *LinkedList) IsPalindrome() bool {
	len := l.length
	switch len {
	case 0:
		return false
	case 1:
		return true
	}
	queue := make([]string, 0, len/2)
	isOdd := len%2 != 0
	cur := l.head.next
	for i := uint(1); i <= len; i++ {
		if isOdd && i == (len/2+1) {
			continue
		}
		if i <= len/2 {
			queue = append(queue, cur.value.(string))
		} else {
			if queue[len-i] != cur.value.(string) {
				return false
			}
		}
		cur = cur.next
	}
	return true
}
// 方法二：带头节点 翻转链表前半部分 时间复杂度O(n) 空间复杂度O(1)
func (l *LinkedList) IsPalindrome2() bool {
	len := l.length
	switch len {
	case 0:
		return false
	case 1:
		return true
	}
	step := len/2
	isPalindrome := true
	//翻转前半部分链表
	var pre *ListNode
	cur := l.head.next
	for i := uint(0); i < step; i ++ {
		cur.next, pre, cur = pre, cur, cur.next
	}
	// 需要记住中间指针以便后续还原
	mid := cur
	var left, right *ListNode = pre, nil
	if len%2 == 0 {
		right = cur
	} else {
		right = cur.next
	}

	for right != nil {
		if left.value.(string) != right.value.(string) {
			isPalindrome = false
			break
		}
		left = left.next
		right = right.next
	}

	// 还原链表
	cur = pre
	pre = mid
	for cur != nil {
		cur.next, pre, cur = pre, cur, cur.next
	}
	return isPalindrome
}
// 方法三：无头节点 快慢指针 时间复杂度O(n) 空间复杂度O(1)
func (l *LinkedList) IsPalindrome3() bool {
	head := l.head.next
	if head == nil {
		return false
	}
	slow, fast := head, head
	var pre, p *ListNode
	for fast != nil && fast.next != nil {
		p = slow
		slow = slow.next
		fast = fast.next.next
		// 翻转链表
		p.next = pre
		pre = p
	}

	left := p // 记录左半链表头，方便还原
	right := slow
	// 奇数，过滤中点
	if fast != nil {
		slow = slow.next
	}

	isPalindrome := true
	for slow != nil {
		if p.value != slow.value {
			isPalindrome = false
			break
		}
		p = p.next
		slow = slow.next
	}

	// 还原链表
	pre = right
	cur := left
	for cur != nil {
		cur.next, pre, cur = pre, cur, cur.next
	}
	l.head.next = pre

	return isPalindrome
}

// 判断链表是否有环
// 方法一(不推荐)：双重遍历 时间复杂度O(n^2) 空间复杂度O(1)
func (l *LinkedList) HasCycle() bool {
	pre := l.head.next
	cur := l.head.next.next
	for cur != nil {
		p := l.head.next
		for p != pre {
			if p == cur {
				return true
			}
			p = p.next
		}
		pre = cur
		cur = cur.next
	}
	return false
}
// 方法二：HashSet存储 时间复杂度O(n) 空间复杂度O(n)
// 方法三：快慢指针 时间复杂度O(n) 空间复杂度O(1)
func (l *LinkedList) HasCycle2() bool {
	slow, fast := l.head.next, l.head.next
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}
	return false
}

// 获取有环链表入环点
// 方法一（不推荐）：双重迭代
func (l *LinkedList) GetEntryNodeOfLoop() *ListNode {
	pre := l.head.next
	cur := l.head.next.next
	for cur != nil {
		p := l.head.next
		for p != pre {
			if p == cur {
				return cur
			}
			p = p.next
		}
		pre = cur
		cur = cur.next
	}
	return nil
}
// 方法二：快慢指针
// 2(L + S) = L + S + nR => L + S = nR => n=1时， L + S = R
func (l *LinkedList) GetEntryNodeOfLoop2() *ListNode {
	slow, fast := l.head.next, l.head.next
	for fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			break
		}
	}
	// 慢指针回到起点
	slow = l.head.next
	for slow != fast {
		slow = slow.next
		fast = fast.next
	}
	return slow
}

// 获取有环链表到入口柄长
func (l *LinkedList) GetLengthToEntry() uint {
	slow, fast := l.head.next, l.head.next
	for fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			break
		}
	}
	// 慢指针回到起点
	len := uint(0)
	slow = l.head.next
	for slow != fast {
		slow = slow.next
		fast = fast.next
		len++
	}
	return len
}

// 获取有环链表环长 快慢指针，相遇后下一次相遇即为环长
func (l *LinkedList) GetLengthOfLoop() uint {
	slow, fast := l.head.next, l.head.next
	for fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			break
		}
	}
	len := uint(0)
	slow = slow.next
	fast = fast.next.next
	for slow != fast {
		slow = slow.next
		fast = fast.next.next
		len++
	}
	return len
}

// 判断两个链表是否相交，以及获取第一个交点
// 分有环和无环
func HasIntersection(l1 *LinkedList, l2 *LinkedList) bool {
	
}
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
