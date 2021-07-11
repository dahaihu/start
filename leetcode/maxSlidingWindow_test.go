package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	//nums := []int{1,3,1,2,0,5}
	//nums := []int{1,3,1,2,0,5}
	k := 3
	nums := []int{7, -8, 7, 5, 7, 1, 6, 0}
	fmt.Println(maxSlidingWindow(nums, k))
	for i := k; i <= len(nums);i++ {
		fmt.Print(maxSlice(nums[i-k:i]))
	}
}

func maxSlice(nums []int) int {
	m := nums[0]
	for _, num :=range nums {
		if num > m {
			m = num
		}
	}
	return m
}