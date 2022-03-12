package leetcode

import "sort"

func medianSlidingWindow(nums []int, k int) []float64 {
	mark := make([]int, k)
	copy(mark, nums[:k])
	sort.Ints(mark)
	result := []float64{mid(mark)}
	for i := k; i < len(nums); i++ {
		insert, remove := nums[i], nums[i-k]
		rmIdx := sort.Search(k, func(j int) bool { return remove <= mark[j] })
		copy(mark[rmIdx:len(mark)-1], mark[rmIdx+1:])
		isIdx := sort.Search(k-1, func(j int) bool { return insert <= mark[j] })
		copy(mark[isIdx+1:k], mark[isIdx:k-1])
		mark[isIdx] = insert
		result = append(result, mid(mark))
	}
	return result
}

func mid(nums []int) float64 {
	midIdx := len(nums) / 2
	if len(nums)%2 == 1 {
		return float64(nums[midIdx])
	}
	return float64(nums[midIdx]+nums[midIdx-1]) / 2
}
