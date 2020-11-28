package leetcode

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	fmt.Println(findSortedArrayTargetLeft([]int{1, 2, 2, 2, 3}, 2))
	fmt.Println(binarySearchLeftPosition([]int{1, 2, 2, 2, 3}, 2))
	fmt.Println(findSortedArrayTargetRight([]int{1, 2, 2, 2, 2, 3}, 2))
}