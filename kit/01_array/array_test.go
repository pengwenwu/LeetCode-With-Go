package array

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	fmt.Println("新建长度为10的数组")
	for i := 0; i < capacity-2; i++ {
		err := arr.InsertToIndex(uint(i), i*i)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
	arr.Print()

	fmt.Println("在第六个插入100")
	err := arr.InsertToIndex(5, 100)
	if nil != err {
		t.Fatal(err.Error())
	}
	arr.Print()

	fmt.Println("尾部插入666")
	err = arr.InsertToTail(666)
	if nil != err {
		t.Fatal(err.Error())
	}
	arr.Print()
}

func TestDelete(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity; i++ {
		err := arr.InsertToIndex(uint(i), i*i)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
	fmt.Println("原数组：")
	arr.Print()
	fmt.Println("删除第一个")
	arr.Delete(0)
	fmt.Println("删除后的数组")
	arr.Print()
}

func TestFind(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity; i++ {
		err := arr.InsertToTail(i * i)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
	arr.Print()

	fmt.Println("查找第6个数")
	fmt.Println(arr.Find(5))
	fmt.Println("查找第12个数")
	fmt.Println(arr.Find(11))
}