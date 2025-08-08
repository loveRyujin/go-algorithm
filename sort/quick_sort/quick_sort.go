package quick_sort

// QuickSort 标准快速排序
// 时间复杂度: 平均 O(n log n)，最坏 O(n^2)
// 空间复杂度: 平均 O(log n)（递归栈）
// 稳定性: 不稳定
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	p := partition(arr, left, right)
	quickSort(arr, left, p-1)
	quickSort(arr, p+1, right)
}

// 采用 Lomuto 分区方案，选择中间元素作为轴并交换到末尾
func partition(arr []int, left, right int) int {
	mid := left + (right-left)/2
	arr[mid], arr[right] = arr[right], arr[mid]
	pivot := arr[right]
	i := left
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return i
}

// QuickSortThreeWay 三路快排（适合重复元素较多的场景）
func QuickSortThreeWay(arr []int) {
	if len(arr) <= 1 {
		return
	}
	quickSortThreeWay(arr, 0, len(arr)-1)
}

func quickSortThreeWay(arr []int, left, right int) {
	if left >= right {
		return
	}
	// 选取中间元素作为轴
	mid := left + (right-left)/2
	pivot := arr[mid]
	lt, i, gt := left, left, right
	for i <= gt {
		if arr[i] < pivot {
			arr[lt], arr[i] = arr[i], arr[lt]
			lt++
			i++
		} else if arr[i] > pivot {
			arr[i], arr[gt] = arr[gt], arr[i]
			gt--
		} else {
			i++
		}
	}
	quickSortThreeWay(arr, left, lt-1)
	quickSortThreeWay(arr, gt+1, right)
}
