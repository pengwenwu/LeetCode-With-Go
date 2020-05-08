package sort

// 冒泡、插入、选择排序
func BubbleSort(arr []int, n int) []int {
	if n <= 1 {
		return arr
	}
	for i := 0; i < n; i++ {
		// 提前结束标识，没有交换则已经全部有序
		flag := false
		for j:= 0; j < n - i - 1; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = tmp
				flag = true
			}
		}
		// 没有交换
		if !flag {
			break
		}
	}
	return arr
}
