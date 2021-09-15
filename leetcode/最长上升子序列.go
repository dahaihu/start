package leetcode

func lengthOfLIS(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	mark := make([]int, length)
	mark[0] = 1
	max := 1
	for idx, num := range nums {
		if idx == 0 {
			continue
		}
		mark[idx] = 1
		for preIdx := idx - 1; preIdx >= 0; preIdx-- {
			if nums[preIdx] >= num {
				continue
			}
			if tmp := mark[preIdx] + 1; tmp > mark[idx] {
				mark[idx] = tmp
			}
		}
		if mark[idx] > max {
			max = mark[idx]
		}
	}
	return max
}

func lengthOfLISBest(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	mark := make([]int, 1)
	mark[0] = nums[0]
	for idx := 1; idx < len(nums); idx++ {
		if nums[idx] > mark[len(mark)-1] {
			mark = append(mark, nums[idx])
			continue
		}
		pos := position(mark, nums[idx])
		mark[pos] = nums[idx]
	}
	return len(mark)
}

func position(nums []int, val int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (right-left)/2 + left
		if val <= nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}


func maxLengthOfLIS(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
