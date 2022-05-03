package leetcode

import (
	"sort"
)

func splitInt(num int) []int {
	var items []int
	for num != 0 {
		items = append(items, num%10)
		num = num / 10
	}
	return items
}

func toInt(items []int) int {
	var result int
	for _, item := range items {
		result = result*10 + item
	}
	return result
}

func curBinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (right-left)/2 + left
		if target <= nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if right == len(nums) {
		return right - 1
	}
	return right
}

func nextNumber(candidates []int, target int) int {
	sort.Ints(candidates)
	targetItems := splitInt(target)
	next := make([]int, 0, len(targetItems))
	var reverseIdx int
	for idx := len(targetItems) - 1; idx >= 0; idx-- {
		cur := targetItems[idx]
		if idx == 0 {
			if cur > candidates[0] {
				ci := curBinarySearch(candidates, cur-1)
				next = append(next, candidates[ci])
			} else {
				reverseIdx = len(next) - 1
				next = append(next, candidates[len(candidates)-1])
				goto reverse
			}
		} else if cur < candidates[0] {
			reverseIdx = len(next) - 1
			for i := idx; i >= 0; i-- {
				next = append(next, candidates[len(candidates)-1])
			}
			goto reverse
		} else {
			ci := curBinarySearch(candidates, cur)
			next = append(next, candidates[ci])
		}
	}
	goto end
reverse:
	for i := reverseIdx; i >= 0; i-- {
		cur := next[i]
		if cur > candidates[0] {
			ci := curBinarySearch(candidates, cur-1)
			next[i] = candidates[ci]
			goto end
		} else {
			next[i] = candidates[len(candidates)-1]
		}
	}
	next = next[1:]
end:
	return toInt(next)
}
