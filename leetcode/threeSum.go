package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] <= nums[j] })
	var result [][]int
	for start := 0; start < len(nums)-2; start++ {
		if start > 0 && nums[start] == nums[start-1] {
			continue
		}
		left, right := start+1, len(nums)-1
		for left < right {
			if cur := nums[start] + nums[left] + nums[right]; cur == 0 {
				result = append(result, []int{nums[start], nums[left], nums[right]})
				left += 1
				for nums[left] == nums[left-1] && left < right {
					left += 1
				}
				right -= 1
				for nums[right] == nums[right+1] && left < right {
					right -= 1
				}
			} else if cur > 0 {
				right -= 1
			} else {
				left += 1
			}
		}
	}
	return result
}
