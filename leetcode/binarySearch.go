package leetcode

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func findSortedArrayTargetLeft(nums []int, target int) int {
	idx := binarySearch(nums, target)
	if idx == -1 {
		return -1
	}
	left, right := 0, idx
	for left < right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func findSortedArrayTargetRight(nums []int, target int) int {
	idx := binarySearch(nums, target)
	if idx == -1 {
		return -1
	}
	left, right := idx, len(nums)-1
	for left < right {
		mid := (right-left)/2 + left + 1
		if nums[mid] == target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

func binarySearchLeftPosition(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (right-left)/2 + left
		if target <= nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
