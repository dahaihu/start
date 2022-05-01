package leetcode

import (
	"fmt"
	"testing"
)

func TestLengthOfLISBest(t *testing.T) {
	nums := []int{0,1,0,3,-1,3}
	fmt.Println("nums is ", nums)
	fmt.Println(lengthOfLISBest(nums))
	nums[len(nums)-1] = -1
	fmt.Println("nums is ", nums)
	fmt.Println(lengthOfLISBest(nums))
}

func TestPosition(t *testing.T) {
	fmt.Println(position([]int{0, 6, 7}, 5))
	fmt.Println(position([]int{1, 3, 3}, 2))
}
