package Stack

import (
	"fmt"
)

// 实现浏览器前进和后退功能
type Browser struct {
	forwardStack Stack
	backStack    Stack
}

func NewBrowser() *Browser {
	return &Browser{
		forwardStack: NewArrayStack(),      // 一个使用数组栈
		backStack:    NewLinkedListSatck(), // 另一个使用链表栈
	}
}

// 判断是否能前进
func (this *Browser) CanForward() bool {
	if this.forwardStack.IsEmpty() {
		return false
	}
	return true
}

// 判断能否后退
func (this *Browser) CanBack() bool {
	if this.backStack.IsEmpty() {
		return false
	}
	return true
}

func (this *Browser) Open(addr string) {
	fmt.Printf("open new addr %+v\n", addr)
	this.forwardStack.Flush()
}

func (this *Browser) PushBack(addr string) {
	this.backStack.Push(addr)
}

func (this *Browser) Forward() {
	if this.forwardStack.IsEmpty() {
		return
	}
	top := this.forwardStack.Pop()
	this.backStack.Push(top)
	fmt.Printf("forward to %+v\n", top)
}

func (this *Browser) Back() {
	if this.backStack.IsEmpty() {
		return
	}
	top := this.backStack.Pop()
	this.forwardStack.Push(top)
	fmt.Printf("back to %+v\n", top)
}
