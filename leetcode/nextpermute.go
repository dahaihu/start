package leetcode

func NextPermutation(nums []int) {
	margin := -1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			margin = i
			break
		}
	}
	if margin == -1 {
		swap(nums, 0, len(nums)-1)
		return
	}

	bigger := -1
	for i := len(nums) - 1; i > margin; i-- {
		if nums[i] > nums[margin] {
			bigger = i
			break
		}
	}
	nums[margin], nums[bigger] = nums[bigger], nums[margin]
	swap(nums, margin+1, len(nums)-1)
}

func swap(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start = start + 1
		end = end - 1
	}
}
