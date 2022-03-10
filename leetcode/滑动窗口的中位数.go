package leetcode

import (
	"sort"
)

func medianSlidingWindow(nums []int, k int) []float64 {
	mark := make([]int, 0, k)
	for i := 0; i < k; i++ {
		mark = append(mark, nums[i])
	}
	sort.Ints(mark)
	result := []float64{mid(mark)}
	for i := k; i < len(nums); i++ {
		rmIdx := sort.Search(k,
			func(j int) bool { return nums[i-k] <= mark[j] })
		for i := rmIdx; i < k-1; i++ {
			mark[i] = mark[i+1]
		}
		adIdx := sort.Search(k-1,
			func(j int) bool { return nums[i] <= mark[j] })
		for j := k - 1; j > adIdx; j-- {
			mark[j] = mark[j-1]
		}
		mark[adIdx] = nums[i]
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
