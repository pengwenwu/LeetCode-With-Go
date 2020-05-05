package recursion

import "fmt"

type FibsMap struct {
	val map[int]int
}

func NewFibs(n int) *FibsMap {
	return &FibsMap{
		val: make(map[int]int, n),
	}
}

func (this *FibsMap) Fibonacci(n int) int {
	if val, ok := this.val[n]; ok {
		return val
	}
	if n <= 1 {
		this.val[1] = 1
		return 1
	} else if n == 2 {
		this.val[2] = 1
		return 1
	} else {
		tmp := this.Fibonacci(n-1) + this.Fibonacci(n-2)
		this.val[n] = tmp
		return tmp
	}
}

func (this *FibsMap) Print(n int) {
	fmt.Println(this.val[n])
}
