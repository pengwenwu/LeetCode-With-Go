package array

import (
	"errors"
	"fmt"
)

/*
通过索引查找数组
在指定索引插入数据
尾部追加数据
删除指定索引数据
打印数组
*/

// 定义
type Array struct {
	data   []int
	length uint
}

// 初始化内存
func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, capacity, capacity),
		length: 0,
	}
}

func (this *Array) Len() uint {
	return this.length
}

// 判定索引是否越界
func (this *Array) isIndexOutOfRange(index uint) bool {
	if index >= this.length {
		return true
	}
	return false
}

// 通过索引查找数组
func (this *Array) Find(index uint) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	return this.data[index], nil
}

// 在指定索引位置插入数据
func (this *Array) InsertToIndex(index uint, v int) error {
	if this.Len() == uint(cap(this.data)) {
		return errors.New("full array")
	}
	if index != this.length && this.isIndexOutOfRange(index) {
		return errors.New("out of range")
	}
	for i := this.length; i > index; i-- {
		this.data[i] = this.data[i-1]
	}
	this.data[index] = v
	this.length++
	return nil
}

// 增加数据
func (this *Array) InsertToTail(v int) error {
	return this.InsertToIndex(this.Len(), v)
}

// 删除指定索引的数据
func (this *Array) Delete(index uint) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	v := this.data[index]
	for i := index; i < this.Len()-1; i++ {
		this.data[i] = this.data[i+1]
	}
	this.length--
	return v, nil
}

// 打印数组
func (this *Array) Print() {
	for i := uint(0); i < this.Len(); i++ {
		fmt.Print(this.data[i], " ")
	}
	fmt.Println()
}
