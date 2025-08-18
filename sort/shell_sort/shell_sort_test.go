package shell_sort

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func init() {
	// 设置随机种子以确保测试结果可重复
	rand.Seed(time.Now().UnixNano())
}

func TestShellSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "正常数组",
			input: []int{64, 34, 25, 12, 22, 11, 90},
			want:  []int{11, 12, 22, 25, 34, 64, 90},
		},
		{
			name:  "已排序数组",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "逆序数组",
			input: []int{5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "包含重复元素",
			input: []int{3, 5, 3, 2, 1, 5},
			want:  []int{1, 2, 3, 3, 5, 5},
		},
		{
			name:  "单个元素",
			input: []int{42},
			want:  []int{42},
		},
		{
			name:  "空数组",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "两个元素",
			input: []int{2, 1},
			want:  []int{1, 2},
		},
		{
			name:  "大量重复元素",
			input: []int{5, 5, 5, 5, 5},
			want:  []int{5, 5, 5, 5, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制输入数组，避免修改原始测试数据
			arr := make([]int, len(tt.input))
			copy(arr, tt.input)

			ShellSort(arr)

			if !reflect.DeepEqual(arr, tt.want) {
				t.Errorf("ShellSort() = %v, want %v", arr, tt.want)
			}
		})
	}
}

func TestShellSortKnuth(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "正常数组",
			input: []int{64, 34, 25, 12, 22, 11, 90},
			want:  []int{11, 12, 22, 25, 34, 64, 90},
		},
		{
			name:  "已排序数组",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "逆序数组",
			input: []int{5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "包含重复元素",
			input: []int{3, 5, 3, 2, 1, 5},
			want:  []int{1, 2, 3, 3, 5, 5},
		},
		{
			name:  "大数组",
			input: []int{9, 1, 8, 2, 7, 3, 6, 4, 5},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制输入数组，避免修改原始测试数据
			arr := make([]int, len(tt.input))
			copy(arr, tt.input)

			ShellSortKnuth(arr)

			if !reflect.DeepEqual(arr, tt.want) {
				t.Errorf("ShellSortKnuth() = %v, want %v", arr, tt.want)
			}
		})
	}
}

func TestShellSortHibbard(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "正常数组",
			input: []int{64, 34, 25, 12, 22, 11, 90},
			want:  []int{11, 12, 22, 25, 34, 64, 90},
		},
		{
			name:  "已排序数组",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "逆序数组",
			input: []int{5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "包含重复元素",
			input: []int{3, 5, 3, 2, 1, 5},
			want:  []int{1, 2, 3, 3, 5, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制输入数组，避免修改原始测试数据
			arr := make([]int, len(tt.input))
			copy(arr, tt.input)

			ShellSortHibbard(arr)

			if !reflect.DeepEqual(arr, tt.want) {
				t.Errorf("ShellSortHibbard() = %v, want %v", arr, tt.want)
			}
		})
	}
}

func TestShellSortSedgewick(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "正常数组",
			input: []int{64, 34, 25, 12, 22, 11, 90},
			want:  []int{11, 12, 22, 25, 34, 64, 90},
		},
		{
			name:  "已排序数组",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "逆序数组",
			input: []int{5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "包含重复元素",
			input: []int{3, 5, 3, 2, 1, 5},
			want:  []int{1, 2, 3, 3, 5, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制输入数组，避免修改原始测试数据
			arr := make([]int, len(tt.input))
			copy(arr, tt.input)

			ShellSortSedgewick(arr)

			if !reflect.DeepEqual(arr, tt.want) {
				t.Errorf("ShellSortSedgewick() = %v, want %v", arr, tt.want)
			}
		})
	}
}

func TestShellSortCiura(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "正常数组",
			input: []int{64, 34, 25, 12, 22, 11, 90},
			want:  []int{11, 12, 22, 25, 34, 64, 90},
		},
		{
			name:  "已排序数组",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "逆序数组",
			input: []int{5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "包含重复元素",
			input: []int{3, 5, 3, 2, 1, 5},
			want:  []int{1, 2, 3, 3, 5, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制输入数组，避免修改原始测试数据
			arr := make([]int, len(tt.input))
			copy(arr, tt.input)

			ShellSortCiura(arr)

			if !reflect.DeepEqual(arr, tt.want) {
				t.Errorf("ShellSortCiura() = %v, want %v", arr, tt.want)
			}
		})
	}
}

func TestShellSortOptimized(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "正常数组",
			input: []int{64, 34, 25, 12, 22, 11, 90},
			want:  []int{11, 12, 22, 25, 34, 64, 90},
		},
		{
			name:  "小数组",
			input: []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
			want:  []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9},
		},
		{
			name:  "很小数组（触发插入排序）",
			input: []int{5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "大数组",
			input: []int{64, 34, 25, 12, 22, 11, 90, 88, 76, 50, 43, 98, 65, 23, 45, 67, 89, 78, 23, 56},
			want:  []int{11, 12, 22, 23, 23, 25, 34, 43, 45, 50, 56, 64, 65, 67, 76, 78, 88, 89, 90, 98},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制输入数组，避免修改原始测试数据
			arr := make([]int, len(tt.input))
			copy(arr, tt.input)

			ShellSortOptimized(arr)

			if !reflect.DeepEqual(arr, tt.want) {
				t.Errorf("ShellSortOptimized() = %v, want %v", arr, tt.want)
			}
		})
	}
}

func TestIsShellSorted(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		gap   int
		want  bool
	}{
		{
			name:  "间隔为1的有序数组",
			input: []int{1, 2, 3, 4, 5},
			gap:   1,
			want:  true,
		},
		{
			name:  "间隔为1的无序数组",
			input: []int{5, 4, 3, 2, 1},
			gap:   1,
			want:  false,
		},
		{
			name:  "间隔为2的有序数组",
			input: []int{1, 3, 2, 4},
			gap:   2,
			want:  true,
		},
		{
			name:  "间隔为2的无序数组",
			input: []int{4, 1, 3, 2}, // 修改测试数据：子序列[4,3]和[1,2]，第一个子序列不有序
			gap:   2,
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsShellSorted(tt.input, tt.gap)
			if result != tt.want {
				t.Errorf("IsShellSorted(%v, %d) = %v, want %v", tt.input, tt.gap, result, tt.want)
			}
		})
	}
}

func TestGetShellSortGaps(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		gapType string
		check   func([]int) bool
	}{
		{
			name:    "Shell间隔序列",
			n:       16,
			gapType: "shell",
			check: func(gaps []int) bool {
				// Shell序列应该是 8, 4, 2, 1
				expected := []int{8, 4, 2, 1}
				return reflect.DeepEqual(gaps, expected)
			},
		},
		{
			name:    "Knuth间隔序列",
			n:       40,
			gapType: "knuth",
			check: func(gaps []int) bool {
				// 检查是否符合Knuth序列模式
				return len(gaps) > 0 && gaps[len(gaps)-1] == 1
			},
		},
		{
			name:    "Hibbard间隔序列",
			n:       20,
			gapType: "hibbard",
			check: func(gaps []int) bool {
				// Hibbard序列应该包含2^k-1形式的数字
				for _, gap := range gaps {
					// 检查gap+1是否是2的幂
					n := gap + 1
					if n <= 0 || (n&(n-1)) != 0 {
						return false
					}
				}
				return true
			},
		},
		{
			name:    "默认序列",
			n:       20,
			gapType: "unknown",
			check: func(gaps []int) bool {
				// 应该返回Knuth序列（默认）
				return len(gaps) > 0
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gaps := GetShellSortGaps(tt.n, tt.gapType)
			if !tt.check(gaps) {
				t.Errorf("GetShellSortGaps(%d, %s) = %v, failed validation", tt.n, tt.gapType, gaps)
			}
		})
	}
}

// 基准测试
func BenchmarkShellSortVariants(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		arr := generateRandomArray(size)

		b.Run(fmt.Sprintf("Shell-size-%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				testArr := make([]int, len(arr))
				copy(testArr, arr)
				ShellSort(testArr)
			}
		})

		b.Run(fmt.Sprintf("Knuth-size-%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				testArr := make([]int, len(arr))
				copy(testArr, arr)
				ShellSortKnuth(testArr)
			}
		})

		b.Run(fmt.Sprintf("Hibbard-size-%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				testArr := make([]int, len(arr))
				copy(testArr, arr)
				ShellSortHibbard(testArr)
			}
		})

		b.Run(fmt.Sprintf("Sedgewick-size-%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				testArr := make([]int, len(arr))
				copy(testArr, arr)
				ShellSortSedgewick(testArr)
			}
		})

		b.Run(fmt.Sprintf("Ciura-size-%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				testArr := make([]int, len(arr))
				copy(testArr, arr)
				ShellSortCiura(testArr)
			}
		})

		b.Run(fmt.Sprintf("Optimized-size-%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				testArr := make([]int, len(arr))
				copy(testArr, arr)
				ShellSortOptimized(testArr)
			}
		})
	}
}

// 性能对比：不同数据分布
func BenchmarkShellSortDataDistribution(b *testing.B) {
	size := 1000

	b.Run("RandomData", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			arr := generateRandomArray(size)
			b.StartTimer()
			ShellSortOptimized(arr)
		}
	})

	b.Run("SortedData", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			arr := generateSortedArray(size)
			b.StartTimer()
			ShellSortOptimized(arr)
		}
	})

	b.Run("ReversedData", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			arr := generateReversedArray(size)
			b.StartTimer()
			ShellSortOptimized(arr)
		}
	})

	b.Run("NearlySortedData", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			arr := generateNearlySortedArray(size)
			b.StartTimer()
			ShellSortOptimized(arr)
		}
	})

	b.Run("DuplicateData", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			arr := generateDuplicateArray(size)
			b.StartTimer()
			ShellSortOptimized(arr)
		}
	})
}

// 不同算法性能对比
func BenchmarkSortingAlgorithmsComparison(b *testing.B) {
	size := 1000
	arr := generateRandomArray(size)

	b.Run("ShellSortOptimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			testArr := make([]int, len(arr))
			copy(testArr, arr)
			ShellSortOptimized(testArr)
		}
	})

	b.Run("GoStandardSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			testArr := make([]int, len(arr))
			copy(testArr, arr)
			sort.Ints(testArr)
		}
	})
}

// 间隔序列性能测试
func BenchmarkGapSequencePerformance(b *testing.B) {
	size := 1000
	arr := generateRandomArray(size)

	sequences := []struct {
		name string
		fn   func([]int)
	}{
		{"Shell", ShellSort},
		{"Knuth", ShellSortKnuth},
		{"Hibbard", ShellSortHibbard},
		{"Sedgewick", ShellSortSedgewick},
		{"Ciura", ShellSortCiura},
		{"Optimized", ShellSortOptimized},
	}

	for _, seq := range sequences {
		b.Run(seq.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				testArr := make([]int, len(arr))
				copy(testArr, arr)
				seq.fn(testArr)
			}
		})
	}
}

// 辅助函数：生成随机数组
func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(1000)
	}
	return arr
}

// 辅助函数：生成已排序数组
func generateSortedArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i
	}
	return arr
}

// 辅助函数：生成逆序数组
func generateReversedArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = size - i - 1
	}
	return arr
}

// 辅助函数：生成近似有序数组
func generateNearlySortedArray(size int) []int {
	arr := generateSortedArray(size)
	// 随机交换少量元素
	swaps := size / 10
	for i := 0; i < swaps; i++ {
		a := rand.Intn(size)
		b := rand.Intn(size)
		arr[a], arr[b] = arr[b], arr[a]
	}
	return arr
}

// 辅助函数：生成重复元素数组
func generateDuplicateArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(10) // 只使用0-9的数字，制造大量重复
	}
	return arr
}
