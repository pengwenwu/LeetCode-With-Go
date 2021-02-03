package _8_sort

import (
    sort "LeetCode-With-Go/kit/07_sort"
)

// 桶排序

// 获取数组中最大的值
func getMax(arr []int) int {
    max := arr[0]
    num := len(arr)
    for i := 1; i < num; i++ {
        if arr[i] > max {
            max = arr[i]
        }
    }
    return max
}

func BucketSort(arr []int)  {
    num := len(arr)
    if num <= 1 {
        return
    }
    maxVal := getMax(arr)
    buckets := make([][]int, num) // 二维切片

    index := 0
    for i := 0; i < num; i++ {
        //index = arr[i] * (num - 1) / maxVal// 桶序号
        index = arr[i] / maxVal * (num - 1) // 桶序号（个人觉得这样更方便理解））
        buckets[index] = append(buckets[index], arr[i]) // 加入到对应的桶中
    }

    tmpPos := 0 // 标记数组位置
    for i := 0; i < num; i++ {
        bucketLen := len(buckets[i])
        if bucketLen > 0 {
            sort.QuickSort(buckets[i]) // 桶内快速排序
            copy(arr[tmpPos:], buckets[i])
            tmpPos += bucketLen
        }
    }
}

// 桶排序简单实现
func BucketSortSimple(source []int)  {
    num := len(source)
    if num <= 1 {
        return
    }

    tmpArr := make([]int, getMax(source)+1)
    for i := 0; i < num; i++ {
        tmpArr[source[i]]++
    }
    c := make([]int, 0)
    for i :=0; i < len(tmpArr); i++ {
        for tmpArr[i] != 0 {
            c = append(c, i)
            tmpArr[i]--
        }
    }
    copy(source, c)
}