package leetcode

import "fmt"

func quickSortSub(nums []int, start, end int) int {
	left, right := start+1, end
	for {
		for left <= right && nums[left] < nums[start] {
			left += 1
		}
		for nums[right] > nums[start] {
			right -= 1
		}
		if left >= right {
			break
		}
		nums[left], nums[right] = nums[right], nums[left]
		left += 1
		right -= 1
	}
	nums[start], nums[right] = nums[right], nums[start]
	return right
}

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	idx := quickSortSub(nums, start, end)
	quickSort(nums, start, idx-1)
	quickSort(nums, idx+1, end)
}


func quickMain(nums []int) {
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
