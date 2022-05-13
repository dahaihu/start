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
	items := splitInt(target)
	suited := -1
	for i := 0; i < len(items); i++ {
		cur := items[i]
		if cur < candidates[len(candidates)-1] {
			biggerIdx := sort.Search(len(candidates), func(i int) bool { return cur < candidates[i] })
			items[i] = candidates[biggerIdx]
			suited = i
			break
		}
	}
	var out []int
	if suited == -1 {
		out = make([]int, len(items)+1)
		for i := 0; i < len(out); i++ {
			out[i] = candidates[0]
		}
	} else {
		for i := 0; i < suited; i++ {
			items[i] = candidates[0]
		}
		out = items
	}
	var res int
	for i := len(out) - 1; i >= 0; i-- {
		res = res*10 + out[i]
	}
	return res
}
