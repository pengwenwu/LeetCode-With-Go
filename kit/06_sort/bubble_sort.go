package sort

// 冒泡、插入、选择排序 时间复杂度O(n^2)
// 冒泡：原地、稳定
func BubbleSort(arr []int, n int) []int {
    if n <= 1 {
        return arr
    }
    for i := 0; i < n; i++ {
        // 提前结束标识，没有交换则已经全部有序
        flag := false
        // 每次循环，最大的数必然在末尾
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                //tmp := arr[j+1]
                //arr[j+1] = arr[j]
                //arr[j] = tmp
                arr[j], arr[j+1] = arr[j+1], arr[j]
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

// 插入排序（原地、稳定）
func InsertionSort(arr []int, n int) []int {
    if n <= 1 {
        return arr
    }
    for i := 1; i < n; i++ {
        value := arr[i]
        j := i - 1
        // 寻找插入的位置
        for ; j >= 0; j-- {
            if arr[j] > value {
                arr[j+1] = arr[j] // 数据移动
            } else {
                break
            }
        }
        arr[j+1] = value // 插入数据
    }
    return arr
}

// 选择排序（原地、非稳定）
func SelectionSort(arr []int, n int) []int {
    if n <= 1 {
        return arr
    }
    for i := 0; i < n; i++ {
        minIndex := i
        for j := i+1; j <= n-1; j++ {
            if arr[j] < arr[minIndex] {
                minIndex = j
            }
        }
        arr[i], arr[minIndex] = arr[minIndex], arr[i]
    }
    return arr
}
