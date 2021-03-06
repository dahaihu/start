package leetcode

import "sort"

// 思考过程，带着过程来进行做题，会舒服很多
// 不要站在一个大的问题上来思考问题，这样会无从下手
// 要思考过程，解决过程中遇到的问题
func PermuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	results := make([][]int, 0)
	var _permute func([]int, []int)
	_permute = func(nums []int, internalResult []int) {
		if len(nums) == 1 {
			internalResult = append(internalResult, nums[0])
			results = append(results, internalResult)
			return
		}
		for i := 0; i < len(nums); i++ {
			if i >= 1 && nums[i] == nums[i-1] {
				continue
			}
			tmpInternalResult := make([]int, len(internalResult))
			copy(tmpInternalResult, internalResult)
			tmpInternalResult = append(tmpInternalResult, nums[i])
			_permute(popItem(nums, i), tmpInternalResult)
		}
	}
	internalResult := make([]int, 0)
	_permute(nums, internalResult)
	return results
}

