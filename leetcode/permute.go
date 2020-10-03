package leetcode

import "sort"

// 思考过程，带着过程来进行做题，会舒服很多
// 不要站在一个大的问题上来思考问题，这样会无从下手
// 要思考过程，解决过程中遇到的问题
func permute(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	var _permute func([]int, []int)
	_permute = func(left []int, tmpRes []int) {
		if len(left) == 1 {
			tmpRes = append(tmpRes, left[0])
			res = append(res, tmpRes)
		}
		for i, val := range left {
			if i >= 1 && val == left[i-1] {
				continue
			}
			_left := popItem(left, i)
			tmp := make([]int, len(tmpRes))
			copy(tmp, tmpRes)
			tmp = append(tmp, val)
			_permute(_left, tmp)
		}
	}
	_permute(nums, make([]int, 0))
	return res

}


func popItem(nums []int, idx int) []int {
	res := make([]int, len(nums)-1)
	copy(res[:idx], nums[:idx])
	copy(res[idx:], nums[idx+1:])
	return res
}
