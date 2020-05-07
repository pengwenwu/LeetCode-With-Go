package sort

// 冒泡、插入、选择排序
func BubbleSort(arr []int, n int) {
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		// 提前结束标识，没有交换则已经全部有序
		flag := false
		for j:= i; j < n; j++ {
			if arr[i] > arr[j] {
				tmp := arr[j]
				arr[j] = arr[i]
				arr[i] = tmp
				flag = true
			}
		}
	}
}
