package leetcode

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) <= 1 {
		return nums
	}
	window := make([]int, 0, k)
	for i := 0; i < k; i++ {
		for len(window) > 0 && window[len(window)-1] < nums[i] {
			window = window[:len(window)-1]
		}
		window = append(window, nums[i])
	}
	result := make([]int, 0, len(nums)-k+1)
	result = append(result, window[0])
	for idx, num := range nums[k:] {
		if len(window) > 0 && window[0] == nums[idx] {
			window = window[1:]
		}
		for len(window) > 0 && window[len(window)-1] < num {
			window = window[:len(window)-1]
		}
		window = append(window, num)
		result = append(result, window[0])
	}
	return result
}
