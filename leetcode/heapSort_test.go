package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func TestHeapSort(t *testing.T) {
	array := []int{2, 1, 3, 5, 6, 6, 4, 7}
	selectHeapSort(array)
	fmt.Println(array)
}

func TestHeapfy(t *testing.T) {
	nums := []int{2, 1, 3, 5, 6, 6, 4, 7}
	h := NewBigHeap(nums)
	fmt.Println(h.nums)
	h.Sort()
	fmt.Println(h.nums)
}

func TestHeap(t *testing.T) {
	nums := []int{2, 1, 3, 5, 6, 6, 4, 7, 5}
	h := NewBigHeap(nums)
	for len(h.nums) != 0 {
		fmt.Println(h.Pop())
	}
}

func TestScheduler(t *testing.T) {
	fmt.Println(leastInterval2(
		[]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'},
		2,
	))
}

func TestSearchBinary(t *testing.T) {
	data := []int{1, 2, 3, 4}
	fmt.Println(sort.Search(len(data), func(i int) bool { return data[i] >= 4 }))
}

func TestChars(t *testing.T) {
	fmt.Println('Z' - 'A')
}

func TestMargin(t *testing.T) {
	fmt.Println([]int{1, 2, 3}[3:])
}
