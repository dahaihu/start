package leetcode

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	// fmt.Println(findSortedArrayTargetLeft([]int{1, 2, 2, 2, 3}, 2))
	fmt.Println(binarySearchLeftPosition([]int{1, 2, 2, 2, 5}, 6))
	// fmt.Println(binarySearchRightPosition([]int{1, 2, 2, 2, 2, 3}, 2))
}
