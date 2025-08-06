package bubble_sort

// BubbleSort 冒泡排序算法实现
// 时间复杂度: O(n²)
// 空间复杂度: O(1)
// 稳定性: 稳定
func BubbleSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// 外层循环控制需要比较的轮数
	for i := 0; i < n-1; i++ {
		// 添加一个标志位，如果某一轮没有交换，说明已经有序
		swapped := false

		// 内层循环进行相邻元素的比较和交换
		// n-i-1 是因为每轮都会把最大的元素放到最后
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// 交换相邻元素
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		// 如果这一轮没有发生交换，说明已经有序，可以提前退出
		if !swapped {
			break
		}
	}
}

// BubbleSortOptimized 优化版冒泡排序
// 记录最后一次交换的位置，该位置之后的元素已经有序
func BubbleSortOptimized(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for n > 1 {
		// 记录最后一次交换的位置
		newN := 0
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				newN = i
			}
		}
		// newN之后的元素已经有序，下一轮只需要比较到newN即可
		n = newN
	}
}
