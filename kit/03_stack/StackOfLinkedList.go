package Stack

import (
	"fmt"
)

type Node struct {
	next *Node
	val  interface{}
}

type LinkedListStack struct {
	topNode *Node // 栈顶节点
}

func NewLinkedListSatck() *LinkedListStack {
	return &LinkedListStack{nil}
}

func (this *LinkedListStack) IsEmpty() bool {
	if this.topNode == nil {
		return true
	}
	return false
}

func (this *LinkedListStack) Push(v interface{}) {
	this.topNode = &Node{next: this.topNode, val: v}
}

func (this *LinkedListStack) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	}
	v := this.topNode.val
	this.topNode = this.topNode.next
	return v
}

func (this *LinkedListStack) Top() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.topNode.val
}

func (this *LinkedListStack) Flush() {
	this.topNode = nil
}

func (this *LinkedListStack) Print() {
	if this.IsEmpty() {
		fmt.Println("emtpy stack")
	} else {
		cur := this.topNode
		for cur != nil {
			fmt.Println(cur.val)
			cur = cur.next
		}
	}
}
