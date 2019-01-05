package Stack

import (
	"fmt"
)

// 基于数组实现的栈
type ArrayStack struct {
	data []interface{} // 数据
	top  int           // 栈顶指针
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

func (this *ArrayStack) IsEmpty() bool {
	if this.top < 0 {
		return true
	}
	return false
}

func (this *ArrayStack) Push(v interface{}) {
	this.top++
	if this.top > len(this.data)-1 {
		this.data = append(this.data, v)
	} else {
		this.data[this.top] = v
	}
}

func (this *ArrayStack) Pop() interface{} {
	if this.top < 0 {
		return nil
	}
	cur := this.data[this.top]
	this.top--
	return cur
}

func (this *ArrayStack) Top() interface{} {
	if this.top < 0 {
		return nil
	}
	return this.data[this.top]
}

func (this *ArrayStack) Flush() {
	this.top = -1
}

func (this *ArrayStack) Print() {
	if this.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		for i := this.top; i >= 0; i-- {
			fmt.Println(this.data[i])
		}
	}
}
