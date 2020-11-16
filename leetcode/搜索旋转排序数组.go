package leetcode

func search(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (right - left)/2 + left
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[0] {
			if target < nums[mid] && target >= nums[0] {
				right = mid
			} else {
				left = mid + 1
			}
		} else {
			if target <= nums[right-1] && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}
	return -1
}
