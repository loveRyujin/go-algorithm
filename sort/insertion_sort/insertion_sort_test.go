package insertion_sort

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"正常数组", []int{64, 34, 25, 12, 22, 11, 90}, []int{11, 12, 22, 25, 34, 64, 90}},
		{"已排序数组", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"逆序数组", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"包含重复元素", []int{3, 5, 3, 2, 1, 5}, []int{1, 2, 3, 3, 5, 5}},
		{"单个元素", []int{42}, []int{42}},
		{"空数组", []int{}, []int{}},
		{"两个元素", []int{2, 1}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := make([]int, len(tt.input))
			copy(arr, tt.input)
			InsertionSort(arr)
			if !reflect.DeepEqual(arr, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", arr, tt.want)
			}
		})
	}
}

func TestBinaryInsertionSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"正常数组", []int{64, 34, 25, 12, 22, 11, 90}, []int{11, 12, 22, 25, 34, 64, 90}},
		{"已排序数组", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"逆序数组", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"包含重复元素", []int{3, 5, 3, 2, 1, 5}, []int{1, 2, 3, 3, 5, 5}},
		{"单个元素", []int{42}, []int{42}},
		{"空数组", []int{}, []int{}},
		{"两个元素", []int{2, 1}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := make([]int, len(tt.input))
			copy(arr, tt.input)
			BinaryInsertionSort(arr)
			if !reflect.DeepEqual(arr, tt.want) {
				t.Errorf("BinaryInsertionSort() = %v, want %v", arr, tt.want)
			}
		})
	}
}

// 基准测试
func BenchmarkInsertionSort(b *testing.B) {
	sizes := []int{10, 100, 1000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := generateRandomArray(size)
				b.StartTimer()
				InsertionSort(arr)
			}
		})
	}
}

func BenchmarkBinaryInsertionSort(b *testing.B) {
	sizes := []int{10, 100, 1000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := generateRandomArray(size)
				b.StartTimer()
				BinaryInsertionSort(arr)
			}
		})
	}
}

func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(1000)
	}
	return arr
}
