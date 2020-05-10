package _7_sort

// 归并排序、快速排序 时间复杂度O(nlogn)
// 分而治之
func MergeSort(arr []int) {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}
	mergeSort(arr, 0, arrLen-1)
}

func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	tmpArr := make([]int, end-start+1)
	i := start
	j := mid + 1
	k := 0
	for ; i <= mid && j <= end; k++ {
		if arr[i] <= arr[j] {
			tmpArr[k] = arr[i]
			i++
		} else {
			tmpArr[k] = arr[j]
			j++
		}
	}
	// 写入剩余数据
	for ; i <= mid; i++ {
		tmpArr[k] = arr[i]
		k++
	}
	for ; j <= end; j++ {
		tmpArr[k] = arr[j]
		k++
	}

	// 覆盖至原有数组指定位置
	copy(arr[start:end+1], tmpArr)
}

// 快速排序
func QuickSort(arr []int) {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}
	separationSort(arr, 0, arrLen-1)
}

func separationSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	// 分区
	q := partition(arr, start, end)
	separationSort(arr, start, q-1)
	separationSort(arr, q+1, end)
}

func partition(arr []int, start, end int) int {
	// 默认选取最后一位为对比数组
	pivot := arr[end]
	i := start
	for j:= start; j < end; j++ {
		if arr[j] >= pivot {
			continue
		}
		if i != j {
			// 交换
			arr[i], arr[j] = arr[j], arr[i]
		}
		i++
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}
