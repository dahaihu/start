package leetcode


func firstMissingPositive(nums []int) int {
	margin := len(nums)+1
	for idx := range nums {
		if nums[idx] < 0 || nums[idx] >= margin {
			nums[idx] = 0
		}
	}

	for _, num := range nums {
		idx := num % margin
		if idx == 0 {
			continue
		}
		nums[idx-1] += margin
	}

	for idx, num := range nums {
		if num < margin {
			return idx+1
		}
	}
	return margin
}
