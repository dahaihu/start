package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	//nums := []int{1,3,1,2,0,5}
	//nums := []int{1,3,1,2,0,5}
	nums := []int{7, -8, 7, 5, 7, 1, 6, 0}
	fmt.Println(maxSlidingWindow(nums, 4))
}