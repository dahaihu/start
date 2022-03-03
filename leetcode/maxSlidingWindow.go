package leetcode

func maxSlidingWindow(nums []int, k int) []int {
	var queue []int
	for i := 0; i < k; i++ {
		for len(queue) != 0 {
			lastIdx := queue[len(queue)-1]
			if nums[i] > nums[lastIdx] {
				queue = queue[:len(queue)-1]
			} else {
				break
			}
		}
		queue = append(queue, i)
	}
	result := []int{nums[queue[0]]}
	for i := k; i < len(nums); i++ {
		if queue[0] <= i-k {
			queue = queue[1:]
		}
		for len(queue) != 0 {
			lastIdx := queue[len(queue)-1]
			if nums[i] > nums[lastIdx] {
				queue = queue[:len(queue)-1]
			} else {
				break
			}
		}
		queue = append(queue, i)
		result = append(result, nums[queue[0]])
	}
	return result
}
