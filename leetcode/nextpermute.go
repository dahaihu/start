package leetcode

func nextPermutation(nums []int) {
	var peekIndex int
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			peekIndex = i + 1
			break
		}
	}
	if peekIndex == 0 {
		reverseArray(nums, 0, len(nums)-1)
		return
	}
	for i := len(nums)-1; i >= peekIndex; i-- {
		if nums[i] > nums[peekIndex-1] {
			nums[i], nums[peekIndex-1] = nums[peekIndex-1], nums[i]
			reverseArray(nums, peekIndex, len(nums)-1)
			return
		}
	}
}

func reverseArray(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
