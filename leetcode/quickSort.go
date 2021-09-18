package leetcode

import "fmt"

func sub(nums []int, start, end int) int {
	left, right := start+1, end
	for {
		for left <= right && nums[left] <= nums[start] {
			left++
		}
		for left <= right && nums[right] >= nums[start] {
			right--
		}
		if left > right {
			break
		}
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
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

func splitArray(nums []int, target int) (int, int) {
	left := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			nums[i], nums[left] = nums[left], nums[i]
			left += 1
		}
	}
	start, end := left, len(nums)-1
	for {
		for start <= end && nums[start] < target {
			start += 1
		}
		for nums[end] > target {
			end -= 1
		}
		if start >= end {
			break
		}
		nums[start], nums[end] = nums[end], nums[start]
		start += 1
		end -= 1
	}

	resRight := end
	for left > 0 {
		nums[left-1], nums[end] = nums[end], nums[left-1]
		left -= 1
		end -= 1
	}
	return end+1, resRight
}


