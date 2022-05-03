package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0, 10)
	for first := 0; first < len(nums)-3; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		for second := first+1; second < len(nums)-2; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			start, end := second+1, len(nums)-1
			for start < end {
				tmp := nums[first] + nums[second] + nums[start] + nums[end]
				if tmp == target {
					res = append(res, []int{nums[first], nums[second], nums[start], nums[end]})
					start += 1
					for start < end && nums[start] == nums[start-1] {
						start += 1
					}
					end -= 1
					for start < end && nums[end] == nums[end+1] {
						end -= 1
					}
				} else if tmp < target {
					start += 1
				} else {
					end -= 1
				}
			}
		}
	}
	return res
}