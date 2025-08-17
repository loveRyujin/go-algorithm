package heap_sort

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestHeapSort(t *testing.T) {
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

			HeapSort(arr)

			if !reflect.DeepEqual(arr, tt.want) {
				t.Errorf("HeapSort() = %v, want %v", arr, tt.want)
			}
		})
	}
}

func TestHeapSortOptimized(t *testing.T) {
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
			name:  "奇数长度数组",
			input: []int{9, 1, 8, 2, 7, 3, 6},
			want:  []int{1, 2, 3, 6, 7, 8, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制输入数组，避免修改原始测试数据
			arr := make([]int, len(tt.input))
			copy(arr, tt.input)

			HeapSortOptimized(arr)

			if !reflect.DeepEqual(arr, tt.want) {
				t.Errorf("HeapSortOptimized() = %v, want %v", arr, tt.want)
			}
		})
	}
}

func TestHeapSortMin(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "正常数组（降序）",
			input: []int{64, 34, 25, 12, 22, 11, 90},
			want:  []int{90, 64, 34, 25, 22, 12, 11},
		},
		{
			name:  "已排序数组（降序）",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{5, 4, 3, 2, 1},
		},
		{
			name:  "包含重复元素（降序）",
			input: []int{3, 5, 3, 2, 1, 5},
			want:  []int{5, 5, 3, 3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制输入数组，避免修改原始测试数据
			arr := make([]int, len(tt.input))
			copy(arr, tt.input)

			HeapSortMin(arr)

			if !reflect.DeepEqual(arr, tt.want) {
				t.Errorf("HeapSortMin() = %v, want %v", arr, tt.want)
			}
		})
	}
}

func TestHeapSortWithCallback(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		less  func(i, j int) bool
		want  []int
		desc  string
	}{
		{
			name:  "升序排列",
			input: []int{3, 1, 4, 1, 5, 9, 2, 6},
			less:  func(i, j int) bool { return i < j },
			want:  []int{1, 1, 2, 3, 4, 5, 6, 9},
			desc:  "使用小于比较函数实现升序",
		},
		{
			name:  "降序排列",
			input: []int{3, 1, 4, 1, 5, 9, 2, 6},
			less:  func(i, j int) bool { return i > j },
			want:  []int{9, 6, 5, 4, 3, 2, 1, 1},
			desc:  "使用大于比较函数实现降序",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制输入数组，避免修改原始测试数据
			arr := make([]int, len(tt.input))
			copy(arr, tt.input)

			HeapSortWithCallback(arr, tt.less)

			if !reflect.DeepEqual(arr, tt.want) {
				t.Errorf("HeapSortWithCallback() = %v, want %v (%s)", arr, tt.want, tt.desc)
			}
		})
	}
}

func TestIsMaxHeap(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  bool
	}{
		{
			name:  "有效的最大堆",
			input: []int{90, 64, 34, 25, 22, 12, 11},
			want:  true,
		},
		{
			name:  "无效的最大堆",
			input: []int{11, 12, 22, 25, 34, 64, 90},
			want:  false,
		},
		{
			name:  "单个元素",
			input: []int{42},
			want:  true,
		},
		{
			name:  "空数组",
			input: []int{},
			want:  true,
		},
		{
			name:  "两个元素（有效）",
			input: []int{2, 1},
			want:  true,
		},
		{
			name:  "两个元素（无效）",
			input: []int{1, 2},
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsMaxHeap(tt.input)
			if result != tt.want {
				t.Errorf("IsMaxHeap(%v) = %v, want %v", tt.input, result, tt.want)
			}
		})
	}
}

func TestIsMinHeap(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  bool
	}{
		{
			name:  "有效的最小堆",
			input: []int{11, 12, 22, 25, 34, 64, 90},
			want:  true,
		},
		{
			name:  "无效的最小堆",
			input: []int{90, 64, 34, 25, 22, 12, 11},
			want:  false,
		},
		{
			name:  "单个元素",
			input: []int{42},
			want:  true,
		},
		{
			name:  "空数组",
			input: []int{},
			want:  true,
		},
		{
			name:  "两个元素（有效）",
			input: []int{1, 2},
			want:  true,
		},
		{
			name:  "两个元素（无效）",
			input: []int{2, 1},
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsMinHeap(tt.input)
			if result != tt.want {
				t.Errorf("IsMinHeap(%v) = %v, want %v", tt.input, result, tt.want)
			}
		})
	}
}

// 测试堆构建过程
func TestBuildMaxHeap(t *testing.T) {
	arr := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	buildMaxHeap(arr)

	if !IsMaxHeap(arr) {
		t.Errorf("buildMaxHeap() failed to create a valid max heap: %v", arr)
	}
}

func TestBuildMinHeap(t *testing.T) {
	arr := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	buildMinHeap(arr)

	if !IsMinHeap(arr) {
		t.Errorf("buildMinHeap() failed to create a valid min heap: %v", arr)
	}
}

// 基准测试
func BenchmarkHeapSort(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := generateRandomArray(size)
				b.StartTimer()

				HeapSort(arr)
			}
		})
	}
}

func BenchmarkHeapSortOptimized(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := generateRandomArray(size)
				b.StartTimer()

				HeapSortOptimized(arr)
			}
		})
	}
}

// 比较不同算法的性能
func BenchmarkHeapSortComparison(b *testing.B) {
	size := 1000
	arr := generateRandomArray(size)

	b.Run("StandardHeapSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			testArr := make([]int, len(arr))
			copy(testArr, arr)
			HeapSort(testArr)
		}
	})

	b.Run("OptimizedHeapSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			testArr := make([]int, len(arr))
			copy(testArr, arr)
			HeapSortOptimized(testArr)
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

// 性能对比：不同数据分布
func BenchmarkHeapSortDataDistribution(b *testing.B) {
	size := 1000

	b.Run("RandomData", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			arr := generateRandomArray(size)
			b.StartTimer()
			HeapSort(arr)
		}
	})

	b.Run("SortedData", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			arr := generateSortedArray(size)
			b.StartTimer()
			HeapSort(arr)
		}
	})

	b.Run("ReversedData", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			arr := generateReversedArray(size)
			b.StartTimer()
			HeapSort(arr)
		}
	})

	b.Run("DuplicateData", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			arr := generateDuplicateArray(size)
			b.StartTimer()
			HeapSort(arr)
		}
	})
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

// 辅助函数：生成重复元素数组
func generateDuplicateArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(10) // 只使用0-9的数字，制造大量重复
	}
	return arr
}
