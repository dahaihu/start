package leetcode

func longestConsecutive(nums []int) int {
	mark := make(map[int]int)
	var result int
	for _, num := range nums {
		if _, ok := mark[num]; ok {
			continue
		}
		leftLen, rightLen := mark[num-1], mark[num+1]
		curLen := leftLen + 1 + rightLen
		if curLen > result {
			result = curLen
		}
		mark[num-leftLen] = curLen
		mark[num+rightLen] = curLen
		mark[num] = curLen
	}
	return result
}
