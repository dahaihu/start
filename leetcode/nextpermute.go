package leetcode

func nextPermutation(nums []int)  {
	margin := len(nums)
	for i:=len(nums)-1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			margin = i
			break
		}
	}
	if margin == len(nums) {
		swap(nums, 0, len(nums)-1)
		return
	}

	for idx := len(nums)-1; idx >= margin; idx-- {
		if nums[idx] > nums[margin-1] {
			nums[margin-1], nums[idx] = nums[idx], nums[margin-1]
			swap(nums, margin, len(nums)-1)
			return
		}
	}
}

func swap(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start += 1
		end -= 1
	}
}


