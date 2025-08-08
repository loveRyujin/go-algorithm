package insertion_sort

// InsertionSort 插入排序
// 时间复杂度: O(n^2)
// 空间复杂度: O(1)
// 稳定性: 稳定
func InsertionSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// BinaryInsertionSort 二分插入排序
// 使用二分查找确定插入位置，减少比较次数
func BinaryInsertionSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		key := arr[i]
		left, right := 0, i-1
		for left <= right {
			mid := left + (right-left)/2
			if arr[mid] > key {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		// 将[left, i-1]区间内元素右移一位
		for j := i - 1; j >= left; j-- {
			arr[j+1] = arr[j]
		}
		arr[left] = key
	}
}
