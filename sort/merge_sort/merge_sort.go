package merge_sort

// MergeSort 标准归并排序
// 时间复杂度: O(n log n)
// 空间复杂度: O(n)
// 稳定性: 稳定
func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	temp := make([]int, len(arr))
	mergeSort(arr, temp, 0, len(arr)-1)
}

func mergeSort(arr, temp []int, left, right int) {
	if left >= right {
		return
	}

	mid := left + (right-left)/2
	mergeSort(arr, temp, left, mid)
	mergeSort(arr, temp, mid+1, right)
	merge(arr, temp, left, mid, right)
}

func merge(arr, temp []int, left, mid, right int) {
	// 复制到临时数组
	for i := left; i <= right; i++ {
		temp[i] = arr[i]
	}

	i, j, k := left, mid+1, left

	// 合并两个有序数组
	for i <= mid && j <= right {
		if temp[i] <= temp[j] {
			arr[k] = temp[i]
			i++
		} else {
			arr[k] = temp[j]
			j++
		}
		k++
	}

	// 复制左边剩余元素
	for i <= mid {
		arr[k] = temp[i]
		i++
		k++
	}

	// 复制右边剩余元素
	for j <= right {
		arr[k] = temp[j]
		j++
		k++
	}
}

// MergeSortBottomUp 自底向上的归并排序（迭代版本）
// 避免递归调用，减少函数调用开销
func MergeSortBottomUp(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	temp := make([]int, n)

	// 从size=1开始，每次将size翻倍
	for size := 1; size < n; size *= 2 {
		// 对每个大小为size的子数组进行归并
		for left := 0; left < n-size; left += 2 * size {
			mid := left + size - 1
			right := min(left+2*size-1, n-1)

			mergeBottomUp(arr, temp, left, mid, right)
		}
	}
}

func mergeBottomUp(arr, temp []int, left, mid, right int) {
	// 复制到临时数组
	for i := left; i <= right; i++ {
		temp[i] = arr[i]
	}

	i, j, k := left, mid+1, left

	// 合并两个有序数组
	for i <= mid && j <= right {
		if temp[i] <= temp[j] {
			arr[k] = temp[i]
			i++
		} else {
			arr[k] = temp[j]
			j++
		}
		k++
	}

	// 复制剩余元素
	for i <= mid {
		arr[k] = temp[i]
		i++
		k++
	}

	for j <= right {
		arr[k] = temp[j]
		j++
		k++
	}
}

// MergeSortInPlace 原地归并排序（优化空间复杂度）
// 使用插入排序的思想实现原地合并，时间复杂度为O(n²)，但空间复杂度为O(1)
func MergeSortInPlace(arr []int) {
	if len(arr) <= 1 {
		return
	}
	mergeSortInPlace(arr, 0, len(arr)-1)
}

func mergeSortInPlace(arr []int, left, right int) {
	if left >= right {
		return
	}

	mid := left + (right-left)/2
	mergeSortInPlace(arr, left, mid)
	mergeSortInPlace(arr, mid+1, right)
	mergeInPlace(arr, left, mid, right)
}

func mergeInPlace(arr []int, left, mid, right int) {
	// 如果已经有序，直接返回
	if arr[mid] <= arr[mid+1] {
		return
	}

	// 使用类似插入排序的方法进行原地合并
	i := left
	j := mid + 1

	for i <= mid && j <= right {
		// 如果左边元素小于等于右边元素，直接跳过
		if arr[i] <= arr[j] {
			i++
		} else {
			// 右边元素更小，需要插入到左边
			value := arr[j]
			index := j

			// 将arr[i..j-1]向右移动一位
			for index != i {
				arr[index] = arr[index-1]
				index--
			}
			arr[i] = value

			// 更新指针
			i++
			mid++
			j++
		}
	}
}

// min 返回两个整数中的较小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
