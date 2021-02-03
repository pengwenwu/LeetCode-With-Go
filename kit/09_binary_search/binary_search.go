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