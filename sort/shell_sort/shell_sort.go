package shell_sort

// ShellSort 标准希尔排序（使用Shell原始间隔序列）
// 时间复杂度: O(n²) 最坏情况，平均情况更好
// 空间复杂度: O(1)
// 稳定性: 不稳定
func ShellSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Shell原始间隔序列: n/2, n/4, n/8, ...
	for gap := n / 2; gap > 0; gap /= 2 {
		// 对每个间隔进行插入排序
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i

			// 在间隔为gap的子序列中进行插入排序
			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = temp
		}
	}
}

// ShellSortKnuth 使用Knuth间隔序列的希尔排序
// Knuth序列: 1, 4, 13, 40, 121, ... (3^k - 1) / 2
// 这个序列在实践中表现很好
func ShellSortKnuth(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// 生成Knuth序列，从最大的开始
	gap := 1
	for gap < n/3 {
		gap = 3*gap + 1
	}

	// 使用Knuth序列进行排序
	for gap >= 1 {
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i

			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = temp
		}
		gap /= 3
	}
}

// ShellSortHibbard 使用Hibbard间隔序列的希尔排序
// Hibbard序列: 1, 3, 7, 15, 31, ... (2^k - 1)
// 理论上时间复杂度为O(n^1.5)
func ShellSortHibbard(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// 生成Hibbard序列
	gaps := generateHibbardSequence(n)

	// 从最大间隔开始排序
	for _, gap := range gaps {
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i

			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = temp
		}
	}
}

// generateHibbardSequence 生成Hibbard间隔序列
func generateHibbardSequence(n int) []int {
	var gaps []int
	k := 1
	for {
		gap := (1 << k) - 1 // 2^k - 1
		if gap >= n {
			break
		}
		gaps = append(gaps, gap)
		k++
	}

	// 反转切片，从大到小排序
	for i, j := 0, len(gaps)-1; i < j; i, j = i+1, j-1 {
		gaps[i], gaps[j] = gaps[j], gaps[i]
	}
	return gaps
}

// ShellSortSedgewick 使用Sedgewick间隔序列的希尔排序
// Sedgewick序列结合了两种模式: 4^i + 3×2^(i-1) + 1 和 4^i - 3×2^i + 1
func ShellSortSedgewick(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// 生成Sedgewick序列
	gaps := generateSedgewickSequence(n)

	// 从最大间隔开始排序
	for _, gap := range gaps {
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i

			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = temp
		}
	}
}

// generateSedgewickSequence 生成Sedgewick间隔序列
func generateSedgewickSequence(n int) []int {
	var gaps []int

	// 生成Sedgewick序列的预定义值
	// 这些是经过计算的Sedgewick序列值
	sedgewickGaps := []int{1, 5, 19, 41, 109, 209, 505, 929, 2161, 3905, 8929, 16001, 36289, 64769, 146305, 260609}

	// 选择小于n的间隔
	for i := len(sedgewickGaps) - 1; i >= 0; i-- {
		if sedgewickGaps[i] < n {
			gaps = append(gaps, sedgewickGaps[i])
		}
	}

	// 如果没有找到合适的间隔，至少包含1
	if len(gaps) == 0 {
		gaps = append(gaps, 1)
	}

	return gaps
}

// ShellSortCiura 使用Ciura间隔序列的希尔排序
// Ciura序列是通过实验得出的经验最优序列
// 序列: 1, 4, 10, 23, 57, 132, 301, 701, 1750, ...
func ShellSortCiura(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// Ciura预定义序列
	ciuraGaps := []int{1750, 701, 301, 132, 57, 23, 10, 4, 1}

	// 找到合适的起始间隔
	var gaps []int
	for _, gap := range ciuraGaps {
		if gap < n {
			gaps = append(gaps, gap)
		}
	}

	// 如果数组很大，扩展Ciura序列
	if len(gaps) == 0 || gaps[0] < n/3 {
		// 使用公式扩展: gap = gap * 2.25
		lastGap := 1750
		for lastGap < n {
			lastGap = int(float64(lastGap) * 2.25)
			if lastGap < n {
				gaps = append([]int{lastGap}, gaps...)
			}
		}
	}

	// 使用选定的间隔进行排序
	for _, gap := range gaps {
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i

			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = temp
		}
	}
}

// ShellSortOptimized 优化版希尔排序
// 结合多种优化策略：
// 1. 对小数组使用插入排序
// 2. 自适应选择最佳间隔序列
// 3. 内循环优化
func ShellSortOptimized(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	// 对很小的数组直接使用插入排序
	if n <= 16 {
		insertionSort(arr)
		return
	}

	// 根据数组大小选择最佳间隔序列
	var gaps []int
	if n <= 1000 {
		// 小数组使用Ciura序列
		gaps = getCiuraGaps(n)
	} else {
		// 大数组使用Knuth序列
		gaps = getKnuthGaps(n)
	}

	// 使用选定的间隔进行排序
	for _, gap := range gaps {
		// 对于gap=1的情况，使用优化的插入排序
		if gap == 1 {
			insertionSortWithGap(arr, gap)
		} else {
			for i := gap; i < n; i++ {
				temp := arr[i]
				j := i

				// 优化：减少数组访问次数
				for j >= gap && arr[j-gap] > temp {
					arr[j] = arr[j-gap]
					j -= gap
				}
				arr[j] = temp
			}
		}
	}
}

// insertionSort 标准插入排序
func insertionSort(arr []int) {
	n := len(arr)
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

// insertionSortWithGap 带间隔的插入排序
func insertionSortWithGap(arr []int, gap int) {
	n := len(arr)
	for i := gap; i < n; i++ {
		temp := arr[i]
		j := i
		for j >= gap && arr[j-gap] > temp {
			arr[j] = arr[j-gap]
			j -= gap
		}
		arr[j] = temp
	}
}

// getCiuraGaps 获取Ciura间隔序列
func getCiuraGaps(n int) []int {
	ciuraBase := []int{1750, 701, 301, 132, 57, 23, 10, 4, 1}
	var gaps []int

	for _, gap := range ciuraBase {
		if gap < n {
			gaps = append(gaps, gap)
		}
	}

	// 如果需要更大的间隔，扩展序列
	if len(gaps) > 0 && gaps[0] < n/3 {
		lastGap := gaps[0]
		for lastGap < n/3 {
			lastGap = int(float64(lastGap) * 2.25)
			gaps = append([]int{lastGap}, gaps...)
		}
	}

	return gaps
}

// getKnuthGaps 获取Knuth间隔序列
func getKnuthGaps(n int) []int {
	var gaps []int
	gap := 1

	// 生成Knuth序列
	for gap < n/3 {
		gap = 3*gap + 1
		gaps = append(gaps, gap)
	}

	// 反转，从大到小
	for i, j := 0, len(gaps)-1; i < j; i, j = i+1, j-1 {
		gaps[i], gaps[j] = gaps[j], gaps[i]
	}

	// 确保最后一个间隔是1
	if len(gaps) == 0 || gaps[len(gaps)-1] != 1 {
		gaps = append(gaps, 1)
	}

	return gaps
}

// IsShellSorted 检查数组是否按间隔排序
func IsShellSorted(arr []int, gap int) bool {
	n := len(arr)
	for i := gap; i < n; i++ {
		if arr[i-gap] > arr[i] {
			return false
		}
	}
	return true
}

// GetShellSortGaps 获取指定类型的间隔序列
func GetShellSortGaps(n int, gapType string) []int {
	switch gapType {
	case "shell":
		return getShellGaps(n)
	case "knuth":
		return getKnuthGaps(n)
	case "hibbard":
		return generateHibbardSequence(n)
	case "sedgewick":
		return generateSedgewickSequence(n)
	case "ciura":
		return getCiuraGaps(n)
	default:
		return getKnuthGaps(n) // 默认使用Knuth序列
	}
}

// getShellGaps 获取Shell原始间隔序列
func getShellGaps(n int) []int {
	var gaps []int
	for gap := n / 2; gap > 0; gap /= 2 {
		gaps = append(gaps, gap)
	}
	return gaps
}
