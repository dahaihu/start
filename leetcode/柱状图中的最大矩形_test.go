package leetcode

import (
	"fmt"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3, 3, 3, 4}) == 14)
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}) == 10)
	fmt.Println(largestRectangleArea([]int{2, 4}) == 4)
	fmt.Println(largestRectangleArea([]int{1}) == 1)
}
