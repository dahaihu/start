package leetcode

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var sameLen int
	result := [][]int{{}}
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			sameLen = len(result)
		}
		curLen := len(result)
		for j := curLen - sameLen; j < curLen; j++ {
			pre := result[j]
			cur := make([]int, len(pre)+1)
			copy(cur[:len(pre)], pre)
			cur[len(pre)] = nums[i]
			result = append(result, cur)
		}
	}
	return result
}
