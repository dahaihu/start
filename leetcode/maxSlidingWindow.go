package leetcode

func maxSlidingWindow(nums []int, k int) []int {
	window := make([]int, 0)
	result := make([]int, 0, len(nums)-k+1)
	for idx, num := range nums {
		if windowBegin := idx - k; windowBegin >= 0 && nums[windowBegin] == window[0] {
			window = window[1:]
		}
		for len(window) != 0 && window[len(window)-1] < num {
			window = window[:len(window)-1]
		}
		window = append(window, num)
		if idx >= k-1 {
			result = append(result, window[0])
		}
	}
	return result
}
