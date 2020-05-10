package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const ArrayLen = 100000

func TestBubbleSort(t *testing.T) {
	var arr []int
	arr = make([]int, ArrayLen)
	rand.Seed(time.Now().Unix())
	for i := 0; i < ArrayLen; i++ {
		arr[i] = rand.Intn(1000)
	}
	startT := time.Now()
	BubbleSort(arr, ArrayLen)
	tc := time.Since(startT)
	fmt.Printf("time cost = %v\n", tc)
}

func TestInsertionSort(t *testing.T) {
	var arr []int
	arr = make([]int, ArrayLen)
	rand.Seed(time.Now().Unix())
	for i := 0; i < ArrayLen; i++ {
		arr[i] = rand.Intn(1000)
	}
	startT := time.Now()
	InsertionSort(arr, ArrayLen)
	tc := time.Since(startT)
	fmt.Printf("time cost = %v\n", tc)
}

func TestSelectionSort(t *testing.T) {
	var arr []int
	arr = make([]int, ArrayLen)
	rand.Seed(time.Now().Unix())
	for i := 0; i < ArrayLen; i++ {
		arr[i] = rand.Intn(1000)
	}
	startT := time.Now()
	SelectionSort(arr, ArrayLen)
	tc := time.Since(startT)
	fmt.Printf("time cost = %v\n", tc)
}