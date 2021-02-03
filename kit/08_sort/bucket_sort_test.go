package _8_sort

import (
    "fmt"
    "math/rand"
    "testing"
    "time"
)

const ArrayLen = 100000
func generateRandArr() []int {
    var arr []int
    arr = make([]int, ArrayLen)
    rand.Seed(time.Now().Unix())
    for i := 0; i < ArrayLen; i++ {
        arr[i] = rand.Intn(1000)
    }
    return arr
}

func TestBucketSort(t *testing.T) {
    arr := generateRandArr()
    //fmt.Println(arr)
    startT := time.Now()
    BucketSort(arr)
    tc := time.Since(startT)
    fmt.Printf("time cost = %v\n", tc)
    //fmt.Println(arr)
}

func TestBucketSortSimple(t *testing.T) {
    arr := generateRandArr()
    //fmt.Println(arr)
    startT := time.Now()
    BucketSortSimple(arr)
    tc := time.Since(startT)
    fmt.Printf("time cost = %v\n", tc)
    //fmt.Println(arr)
}