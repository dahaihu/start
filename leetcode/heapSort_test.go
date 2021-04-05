package leetcode

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	array := []int{2, 1, 3, 5, 6, 6, 4, 7}
	heapSort(array)
	fmt.Println(array)
}
