package leetcode

func lengthOfLIS(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	res := 0
	mark := make([]int, length)
	mark[0] = 1
	for idx := 1; idx < length; idx++ {
		mark[idx] = 1
		height := nums[idx]
		for preIdx := 0; preIdx < idx; preIdx++ {
			if height > nums[preIdx] {
				mark[idx] = maxLengthOfLIS(mark[idx], mark[preIdx]+1)
			}
		}
		res = maxLengthOfLIS(mark[idx], res)
	}
	return res
}

func lengthOfLISBest(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	mark := make([]int, 1)
	mark[0] = nums[0]
	for idx := 1; idx < length; idx++ {
		pos := position(mark, nums[idx])
		if pos == len(mark) {
			mark = append(mark, nums[idx])
		} else {
			mark[pos] = min(mark[pos], nums[idx])
		}
	}
	return len(mark)
}

func position(nums []int, val int) (idx int) {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > val {
			right = mid-1
		} else if nums[mid] < val {
			left = mid+1
		} else {
			return left
		}
	}
	return left
}

func maxLengthOfLIS(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
