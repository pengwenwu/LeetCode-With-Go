package _7_sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const ArrayLen = 100000
func Test_mergeSort(t *testing.T) {
	var arr []int
	arr = make([]int, ArrayLen)
	rand.Seed(time.Now().Unix())
	for i := 0; i < ArrayLen; i++ {
		arr[i] = rand.Intn(1000)
	}
	startT := time.Now()
	MergeSort(arr)
	tc := time.Since(startT)
	fmt.Printf("time cost = %v\n", tc)
}

func TestQuickSort(t *testing.T) {
	var arr []int
	arr = make([]int, ArrayLen)
	rand.Seed(time.Now().Unix())
	for i := 0; i < ArrayLen; i++ {
		arr[i] = rand.Intn(1000)
	}
	startT := time.Now()
	QuickSort(arr)
	tc := time.Since(startT)
	fmt.Printf("time cost = %v\n", tc)
}