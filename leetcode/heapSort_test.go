package leetcode

import (
	"fmt"
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
	fmt.Println(leastIntervalBetter(
		[]byte{'A', 'A', 'A', 'E', 'F', 'G'},
		2,
	))
}

func TestChars(t *testing.T) {
	fmt.Println('Z' - 'A')
}