package leetcode

import (
	"math"
	"sort"
)

// how to initiate
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	sort.Slice(coins, func(i, j int) bool { return coins[i] <= coins[j] })
	mark := make([]int, amount+1)
	for curAmount := coins[0]; curAmount <= amount; curAmount++ {
		min := math.MaxInt64
		for _, coin := range coins {
			if curAmount < coin {
				break
			}
			if (curAmount == coin || mark[curAmount-coin] > 0) && mark[curAmount-coin] < min {
				min = mark[curAmount-coin]
			}
		}
		if min == math.MaxInt64 {
			continue
		}
		mark[curAmount] = min + 1
	}
	if mark[amount] == 0 {
		return -1
	}
	return mark[amount]
}
