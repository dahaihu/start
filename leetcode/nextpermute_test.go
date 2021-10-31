package leetcode

import (
	"fmt"
	"testing"
)

func TestNextPermute(t *testing.T) {
	nums := []int{2, 2, 0, 4, 3, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}
