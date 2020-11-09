package leetcode

import (
	"math"
)

/**
* @Author: 胡大海
* @Date: 2020-03-20 09:02
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */
// K 代表鸡蛋的个数， N 代表楼层的高度
func SuperEggDrop(K int, N int) [][]int {
	mark := make([][]int, N)
	for i := 0; i < N; i++ {
		mark[i] = make([]int, K)
		mark[i][0] = i + 1
	}
	for i := 0; i < K; i++ {
		mark[0][i] = 1
	}
	for floor := 1; floor < N; floor++ {
		for eggs := 1; eggs < K; eggs++ {
			val := math.MaxInt64
			for k := 0; k <= floor; k++ {
				var breakCondition, noBreakCondition int
				if k == 0 {
					breakCondition = 0
				} else {
					breakCondition = mark[k-1][eggs-1]
				}

				if k == floor {
					noBreakCondition = 0
				} else {
					noBreakCondition = mark[floor-k-1][eggs]
				}
				val = min(val, max(breakCondition, noBreakCondition))
			}
			mark[floor][eggs] = val + 1
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
