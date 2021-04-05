package leetcode

func search(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return mid
		} else if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right-1] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}
	return -1
}
