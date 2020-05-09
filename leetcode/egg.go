package leetcode

/**
* @Author: 胡大海
* @Date: 2020-03-20 09:02
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

func SuperEggDrop(K int, N int) [][]int {
	mark := make([][]int, K)
	for idx := 0; idx < K; idx++ {
		mark[idx] = make([]int, N)
	}
	for idx := 0; idx < K; idx++ {
		mark[idx][0] = 1
	}
	for idx := 0; idx < N; idx++ {
		mark[0][idx] = idx + 1
	}

	for times := 2; times <= K; times++ {
		for floors := 2; floors <= N; floors++ {
			minVal := 2 << 30
			for i := 1; i <= floors; i++ {
				var breakCondition, noBreakCondition int
				if i == 1 {
					breakCondition = 0
				} else {
					breakCondition = mark[times-1-1][i-1-1]
				}
				if i == floors {
					noBreakCondition = 0
				} else {
					noBreakCondition = mark[times-1][floors-i-1]
				}
				tmpVal := max(breakCondition, noBreakCondition)
				if tmpVal < minVal {
					minVal = tmpVal
				}
			}
			mark[times-1][floors-1] = minVal + 1
		}
	}
	return mark
}

func min(arrays []int) int{
	minVal := arrays[0]
	for idx := 1; idx < len(arrays); idx++{
		if arrays[idx] < minVal {
			minVal = arrays[idx]
		}
	}
	return minVal
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
