package bubble_sort

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制输入数组，避免修改原始测试数据
			arr := make([]int, len(tt.input))
			copy(arr, tt.input)

			BubbleSort(arr)

			if !reflect.DeepEqual(arr, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", arr, tt.want)
			}
		})
	}
}

func TestBubbleSortOptimized(t *testing.T) {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制输入数组，避免修改原始测试数据
			arr := make([]int, len(tt.input))
			copy(arr, tt.input)

			BubbleSortOptimized(arr)

			if !reflect.DeepEqual(arr, tt.want) {
				t.Errorf("BubbleSortOptimized() = %v, want %v", arr, tt.want)
			}
		})
	}
}

// 基准测试
func BenchmarkBubbleSort(b *testing.B) {
	sizes := []int{10, 100, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := generateRandomArray(size)
				b.StartTimer()

				BubbleSort(arr)
			}
		})
	}
}

func BenchmarkBubbleSortOptimized(b *testing.B) {
	sizes := []int{10, 100, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := generateRandomArray(size)
				b.StartTimer()

				BubbleSortOptimized(arr)
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
