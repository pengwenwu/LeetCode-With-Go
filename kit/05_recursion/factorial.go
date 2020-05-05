package recursion

import "fmt"

// 阶乘
type FactorialMap struct {
	val map[int]int
}

func NewFactorial(n int) *FactorialMap {
	return &FactorialMap{
		val:make(map[int]int, n),
	}
}

func (this *FactorialMap) Factorial(n int) int {
	if val, ok := this.val[n]; ok {
		return val
	}
	if n <= 1 {
		this.val[n] = 1
		return 1
	}
	tmp := n * this.Factorial(n - 1)
	this.val[n] = tmp
	return tmp
}

func (this *FactorialMap) Print(n int) {
	fmt.Println(this.val[n])
}