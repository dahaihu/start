package leetcode

import "fmt"

func sub(nums []int, start, end int) int {
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
	nums[right], nums[start] = nums[start], nums[right]
	return right
}

func quickSort(nums []int, start, end int) {
	if start < end {
		idx := sub(nums, start, end)
		quickSort(nums, start, idx-1)
		quickSort(nums, idx+1, end)
	}
}


func quickMain(nums []int) {
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
