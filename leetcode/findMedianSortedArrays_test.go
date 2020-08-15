package leetcode

import (
	"fmt"
	"testing"
)

func TestFindKth(t *testing.T) {
	// how to get out of the closed quan
	fmt.Println(findKth([]int{1}, []int{2, 3, 4, 5, 6}, 0, 0, 0, 4, 6))
}

func TestDividedSortedArrays(t *testing.T) {
	fmt.Println(dividedTwoSortedArrays([]int{1}, []int{2,3,4,5}))
}
