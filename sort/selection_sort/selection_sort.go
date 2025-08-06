package selection_sort

// SelectionSort 选择排序算法实现
// 时间复杂度: O(n²)
// 空间复杂度: O(1)
// 稳定性: 不稳定
func SelectionSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// 外层循环，每次选择未排序部分的最小元素
	for i := 0; i < n-1; i++ {
		// 假设当前位置的元素是最小的
		minIndex := i

		// 在未排序部分找到最小元素的索引
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		// 如果找到了更小的元素，则交换
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}

// SelectionSortOptimized 优化版选择排序
// 同时找到最大值和最小值，减少交换次数
func SelectionSortOptimized(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	left := 0
	right := n - 1

	for left < right {
		minIndex := left
		maxIndex := right

		// 找到当前范围内的最小值和最大值
		for i := left; i <= right; i++ {
			if arr[i] < arr[minIndex] {
				minIndex = i
			}
			if arr[i] > arr[maxIndex] {
				maxIndex = i
			}
		}

		// 将最小值放到左边
		if minIndex != left {
			arr[left], arr[minIndex] = arr[minIndex], arr[left]
		}

		// 如果最大值在左边，需要调整位置
		if maxIndex == left {
			maxIndex = minIndex
		}

		// 将最大值放到右边
		if maxIndex != right {
			arr[right], arr[maxIndex] = arr[maxIndex], arr[right]
		}

		left++
		right--
	}
}
