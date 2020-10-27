package leetcode

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		// maybe mid == left, then recursive
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if target <= nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target <= nums[right] && target >= nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
