package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func Test_medianSlidingWindow(t *testing.T) {
	fmt.Println(medianSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

func Test_insertIndex(t *testing.T) {
	nums := []int{-3, 1, 3}
	fmt.Println(sort.Search(len(nums), func(i int) bool { return nums[i] >= 3 }))
	fmt.Println(sort.Search(len(nums), func(i int) bool { return nums[i] >= 5 }))
}
