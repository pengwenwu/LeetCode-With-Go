package _9_binary_search

func BinarySearch(arr []int, findVal int) int {
    num := len(arr)
    if num == 0 {
        return -1
    }

    low := 0
    high := num - 1
    for low <= high {
        mid := (low + high) / 2
        if arr[mid] == findVal {
            return mid
        } else if arr[mid] > findVal {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }

    return -1
}

func BinarySearchRecursive(arr []int, findVal int) int {
    num := len(arr)
    if num == 0 {
        return -1
    }

    return bs(arr, findVal, 0, num-1)
}

func bs(arr []int, findVal, low ,high int) int {
    if low > high {
        return -1
    }
    mid := (low + high) / 2
    if arr[mid] == findVal {
        return mid
    } else if arr[mid] > findVal {
        return bs(arr, findVal, low, mid-1)
    } else {
        return bs(arr, findVal, mid+1, high)
    }
}

// 二分查找第一个值等于给定值的元素
func BinarySearchFirst(arr []int, findVal int) int {
    num := len(arr)
    if num == 0 {
        return -1
    }
    low := 0
    high := num - 1
    for low <= high {
        mid := (low + high) / 2
        if arr[mid] >= findVal {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }
    if low < num && arr[low] == findVal {
        return low
    } else {
        return -1
    }
    
    return 0
}

func BinarySearchFirstSimple(arr []int, findVal int) int {
    num := len(arr)
    if num == 0 {
        return -1
    }

    low := 0
    high := num - 1
    for low <= high {
        mid := (low + high) / 2
        if arr[mid] > findVal {
            high = mid - 1
        } else if arr[mid] < findVal {
            low = mid + 1
        } else {
            if mid == 0 || arr[mid-1] != findVal {
                return mid
            } else {
                high = mid - 1
            }
        }
    }
    return -1
}

func BinarySearchLast(arr []int, findVal int) int {
    num := len(arr)
    if num == 0 {
        return -1
    }

    low := 0
    high := num - 1
    for low <= high {
        mid := (low + high) / 2
        if arr[mid] > findVal {
            high = mid - 1
        } else if arr[mid] < findVal {
            low = mid + 1
        } else {
            if mid == num - 1 || arr[mid+1] != findVal {
                return mid
            } else {
                low = mid + 1
            }
        }
    }
    return -1
}

// 查找第一个大于等于给定值的元素
func BinarySearchFirstGT(arr []int, findVal int) int {
    num := len(arr)
    if num == 0 {
        return -1
    }

    low := 0
    high := num - 1
    for low <= high {
        mid := low + (high - low) / 2
        if arr[mid] >= findVal {
            if mid == 0 || arr[mid-1] < findVal {
                return mid
            } else {
                high = mid - 1
            }
        } else {
            low = mid + 1
        }
    }
    return -1
}

// 查找最后一个小于等于给定值的元素
func BinarySearchLastLT(arr []int, findVal int) int {
    num := len(arr)
    if num == 0 {
        return -1
    }

    low := 0
    high := num - 1
    for low <= high {
        mid := low + (high - low) / 2
        if arr[mid] > findVal {
            high = mid - 1
        } else {
            if mid == num - 1 || arr[mid+1] > findVal {
                return mid
            } else {
                low = mid + 1
            }
        }
    }

    return -1
}