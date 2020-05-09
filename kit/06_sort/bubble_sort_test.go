package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestBubbleSort(t *testing.T) {
	const n = 10000
	var arr []int
	arr = make([]int, n)
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(1000)
	}
	startT := time.Now()
	BubbleSort(arr, n)
	tc := time.Since(startT)
	fmt.Printf("time cost = %v\n", tc)
}

func TestInsertionSort(t *testing.T) {
	const n = 10000
	var arr []int
	arr = make([]int, n)
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(1000)
	}
	startT := time.Now()
	InsertionSort(arr, n)
	tc := time.Since(startT)
	fmt.Printf("time cost = %v\n", tc)
}

func TestSelectionSort(t *testing.T) {
	const n = 10000
	var arr []int
	arr = make([]int, n)
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(1000)
	}
	startT := time.Now()
	SelectionSort(arr, n)
	tc := time.Since(startT)
	fmt.Printf("time cost = %v\n", tc)
}