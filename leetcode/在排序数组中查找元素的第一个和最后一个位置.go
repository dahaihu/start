package leetcode

func searchRange(nums []int, target int) []int {
	leftMargin := searchLeftMargin(nums, target)
	if leftMargin >= 0 && leftMargin <= len(nums)-1 && nums[leftMargin] == target {
		return []int{leftMargin, searchLeftMargin(nums, target+1)-1}
	}
	return []int{-1, -1}
}


func searchLeftMargin(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (right-left)/2 + left
		if target <= nums[mid] {
			right = mid
		} else {
			left = mid+1
		}
	}
	return left // 这个地方是不可以判断下是否存在的
}
