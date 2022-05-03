package leetcode

import (
	"math"
)

func canJump(nums []int) bool {
	margin := 0
	for idx, num := range nums {
		if idx > margin {
			return false
		}
		if next := idx + num; next > margin {
			margin = next
		}
	}
	return true
}

func jump(nums []int) int {
	steps := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		steps[i] = math.MaxInt64
	}
	for i := 0; i < len(nums); i++ {
		next := i + nums[i]
		nextStep := steps[i] + 1
		for j := i + 1; j <= next; j++ {
			if j > len(nums)-1 {
				break
			}
			if nextStep < steps[j] {
				steps[j] = nextStep
			}
		}
	}
	return steps[len(steps)-1]
}
