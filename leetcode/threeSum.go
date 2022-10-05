package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] <= nums[j] })
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			cur := nums[i] + nums[left] + nums[right]
			switch {
			case cur == 0:
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left += 1
				for left < right && nums[left] == nums[left-1] {
					left += 1
				}
				right -= 1
				for left < right && nums[right] == nums[right+1] {
					right -= 1
				}
			case cur > 0:
				right -= 1
			case cur < 0:
				left += 1
			}
		}
	}
	return result
}
