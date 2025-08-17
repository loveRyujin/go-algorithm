package heap_sort

// HeapSort 标准堆排序（最大堆）
// 时间复杂度: O(n log n)
// 空间复杂度: O(1)
// 稳定性: 不稳定
func HeapSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// 构建最大堆
	buildMaxHeap(arr)

	// 依次将堆顶元素（最大值）与最后一个元素交换
	// 然后对剩余元素重新构建最大堆
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		maxHeapify(arr, 0, i)
	}
}

// buildMaxHeap 构建最大堆
func buildMaxHeap(arr []int) {
	n := len(arr)
	// 从最后一个非叶子节点开始，向上调整
	for i := n/2 - 1; i >= 0; i-- {
		maxHeapify(arr, i, n)
	}
}

// maxHeapify 维护最大堆性质
// 确保以i为根的子树满足最大堆性质
func maxHeapify(arr []int, i, heapSize int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	// 找到i、left、right中的最大值
	if left < heapSize && arr[left] > arr[largest] {
		largest = left
	}

	if right < heapSize && arr[right] > arr[largest] {
		largest = right
	}

	// 如果最大值不是根节点，则交换并递归调整
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		maxHeapify(arr, largest, heapSize)
	}
}

// HeapSortOptimized 优化版堆排序
// 使用Floyd建堆算法，减少比较次数
func HeapSortOptimized(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// 使用Floyd算法构建最大堆
	buildMaxHeapFloyd(arr)

	// 排序阶段
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		maxHeapifyIterative(arr, 0, i)
	}
}

// buildMaxHeapFloyd 使用Floyd算法构建最大堆
// 自底向上构建，减少比较次数
func buildMaxHeapFloyd(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		siftDown(arr, i, n)
	}
}

// siftDown Floyd算法的下沉操作
func siftDown(arr []int, start, end int) {
	root := start
	for 2*root+1 < end {
		child := 2*root + 1
		swap := root

		// 找到根节点和其子节点中的最大值
		if arr[swap] < arr[child] {
			swap = child
		}
		if child+1 < end && arr[swap] < arr[child+1] {
			swap = child + 1
		}

		if swap == root {
			return
		}

		arr[root], arr[swap] = arr[swap], arr[root]
		root = swap
	}
}

// maxHeapifyIterative 迭代版本的最大堆调整
// 避免递归调用，节省栈空间
func maxHeapifyIterative(arr []int, i, heapSize int) {
	for {
		largest := i
		left := 2*i + 1
		right := 2*i + 2

		if left < heapSize && arr[left] > arr[largest] {
			largest = left
		}

		if right < heapSize && arr[right] > arr[largest] {
			largest = right
		}

		if largest == i {
			break
		}

		arr[i], arr[largest] = arr[largest], arr[i]
		i = largest
	}
}

// HeapSortMin 使用最小堆的堆排序（降序排列）
func HeapSortMin(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// 构建最小堆
	buildMinHeap(arr)

	// 依次将堆顶元素（最小值）与最后一个元素交换
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		minHeapify(arr, 0, i)
	}
}

// buildMinHeap 构建最小堆
func buildMinHeap(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		minHeapify(arr, i, n)
	}
}

// minHeapify 维护最小堆性质
func minHeapify(arr []int, i, heapSize int) {
	smallest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < heapSize && arr[left] < arr[smallest] {
		smallest = left
	}

	if right < heapSize && arr[right] < arr[smallest] {
		smallest = right
	}

	if smallest != i {
		arr[i], arr[smallest] = arr[smallest], arr[i]
		minHeapify(arr, smallest, heapSize)
	}
}

// HeapSortWithCallback 带回调函数的堆排序
// 可以自定义比较函数，支持任意类型的排序
func HeapSortWithCallback(arr []int, less func(i, j int) bool) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// 构建堆
	buildHeapWithCallback(arr, less)

	// 排序
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapifyWithCallback(arr, 0, i, less)
	}
}

// buildHeapWithCallback 使用自定义比较函数构建堆
func buildHeapWithCallback(arr []int, less func(i, j int) bool) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapifyWithCallback(arr, i, n, less)
	}
}

// heapifyWithCallback 使用自定义比较函数维护堆性质
func heapifyWithCallback(arr []int, i, heapSize int, less func(i, j int) bool) {
	target := i
	left := 2*i + 1
	right := 2*i + 2

	// 根据比较函数找到目标节点
	// 如果left比target大（根据less函数），则选择left
	if left < heapSize && less(arr[target], arr[left]) {
		target = left
	}

	// 如果right比target大（根据less函数），则选择right
	if right < heapSize && less(arr[target], arr[right]) {
		target = right
	}

	if target != i {
		arr[i], arr[target] = arr[target], arr[i]
		heapifyWithCallback(arr, target, heapSize, less)
	}
}

// IsMaxHeap 检查数组是否满足最大堆性质
func IsMaxHeap(arr []int) bool {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		left := 2*i + 1
		right := 2*i + 2

		if left < n && arr[i] < arr[left] {
			return false
		}
		if right < n && arr[i] < arr[right] {
			return false
		}
	}
	return true
}

// IsMinHeap 检查数组是否满足最小堆性质
func IsMinHeap(arr []int) bool {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		left := 2*i + 1
		right := 2*i + 2

		if left < n && arr[i] > arr[left] {
			return false
		}
		if right < n && arr[i] > arr[right] {
			return false
		}
	}
	return true
}
