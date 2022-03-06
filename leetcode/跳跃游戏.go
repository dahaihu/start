package leetcode

import (
	"math"
)

func canJump(nums []int) bool {
	maxPos := 0
	for i := 0; i < len(nums); i++ {
		if i > maxPos {
			return false
		}
		if next := i + nums[i]; next > maxPos {
			maxPos = next
		}
		if maxPos >= len(nums)-1 {
			return true
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
