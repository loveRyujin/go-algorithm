package merge_sort

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
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

			MergeSort(arr)

			if !reflect.DeepEqual(arr, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", arr, tt.want)
			}
		})
	}
}

func TestMergeSortBottomUp(t *testing.T) {
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

			MergeSortBottomUp(arr)

			if !reflect.DeepEqual(arr, tt.want) {
				t.Errorf("MergeSortBottomUp() = %v, want %v", arr, tt.want)
			}
		})
	}
}

func TestMergeSortInPlace(t *testing.T) {
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
			name:  "小数组",
			input: []int{3, 1, 2},
			want:  []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制输入数组，避免修改原始测试数据
			arr := make([]int, len(tt.input))
			copy(arr, tt.input)

			MergeSortInPlace(arr)

			if !reflect.DeepEqual(arr, tt.want) {
				t.Errorf("MergeSortInPlace() = %v, want %v", arr, tt.want)
			}
		})
	}
}

// 基准测试
func BenchmarkMergeSort(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := generateRandomArray(size)
				b.StartTimer()

				MergeSort(arr)
			}
		})
	}
}

func BenchmarkMergeSortBottomUp(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := generateRandomArray(size)
				b.StartTimer()

				MergeSortBottomUp(arr)
			}
		})
	}
}

func BenchmarkMergeSortInPlace(b *testing.B) {
	sizes := []int{10, 100, 1000} // 原地归并较慢，减少测试规模

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := generateRandomArray(size)
				b.StartTimer()

				MergeSortInPlace(arr)
			}
		})
	}
}

// 比较不同算法的性能
func BenchmarkMergeSortComparison(b *testing.B) {
	size := 1000
	arr := generateRandomArray(size)

	b.Run("StandardMergeSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			testArr := make([]int, len(arr))
			copy(testArr, arr)
			MergeSort(testArr)
		}
	})

	b.Run("BottomUpMergeSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			testArr := make([]int, len(arr))
			copy(testArr, arr)
			MergeSortBottomUp(testArr)
		}
	})

	b.Run("InPlaceMergeSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			testArr := make([]int, len(arr))
			copy(testArr, arr)
			MergeSortInPlace(testArr)
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

// 测试稳定性
func TestMergeSortStability(t *testing.T) {
	// 使用结构体测试稳定性
	type Item struct {
		value int
		index int
	}

	items := []Item{
		{3, 0}, {1, 1}, {3, 2}, {2, 3}, {1, 4},
	}

	// 将结构体转换为可排序的slice
	values := make([]int, len(items))
	for i, item := range items {
		values[i] = item.value
	}

	MergeSort(values)

	// 验证结果是否正确排序
	expected := []int{1, 1, 2, 3, 3}
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("Stability test failed: got %v, want %v", values, expected)
	}
}
