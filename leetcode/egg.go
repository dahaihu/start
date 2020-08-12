package leetcode

/**
* @Author: 胡大海
* @Date: 2020-03-20 09:02
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */
// K 代表鸡蛋的个数， N 代表楼层的高度
func SuperEggDrop(K int, N int) [][]int {
	mark := make([][]int, K)
	// 初始化二位数组
	for i := 0; i < K; i++ {
		mark[i] = make([]int, N)
		// 在楼层是 1 的时候，次数都是1
		mark[i][0] = 1
	}
	// 在只有一个鸡蛋的时候，只能从一楼不断的往上尝试
	for i := 0; i < N; i++ {
		mark[0][i] = i + 1
	}
	for egg := 2; egg <= K; egg++ {
		for floor := 2; floor <= N; floor++ {
			var noBreakCondition, breakCondition int
			tmpValue := 100 * 10000
			for i := 1; i <= floor; i++ {
				if i == 1 {
					breakCondition = 0
				} else {
					breakCondition = mark[egg-1-1][i-1-1]
				}

				if i == floor {
					noBreakCondition = 0
				} else {
					noBreakCondition = mark[egg-1][floor-i-1]
				}

				maxTmp := max(breakCondition, noBreakCondition)
				if maxTmp < tmpValue {
					tmpValue = maxTmp
				}
			}
			mark[egg-1][floor-1] = tmpValue + 1
		}
	}

	return mark
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
