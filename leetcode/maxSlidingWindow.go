package leetcode

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) <= 1 {
		return nums
	}
	window := make([]int, 0, k)
	push := func(index int) {
		for len(window) > 0 && nums[window[len(window)-1]] < nums[index] {
			window = window[:len(window)-1]
		}
		window = append(window, index)
	}
	for i := 0; i < k; i++ {
		push(i)
	}
	result := make([]int, 0, len(nums)-k+1)
	result = append(result, nums[window[0]])
	for i := k; i < len(nums); i++ {
		if window[0] == i-k {
			window = window[1:]
		}
		push(i)
		result = append(result, nums[window[0]])
	}
	return result
}
