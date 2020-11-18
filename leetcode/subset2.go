package leetcode

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{{}}
	var l int
	for idx, num := range nums {
		length := len(res)
		if idx == 0 || nums[idx] != nums[idx-1] {
			l = len(res)
		}
		for i := length - l; i < length; i++ {
			tmp := make([]int, len(res[i])+1)
			copy(tmp[:len(res[i])], res[i])
			tmp[len(res[i])] = num
			res = append(res, tmp)
		}
	}
	return res
}
