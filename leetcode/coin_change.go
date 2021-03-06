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
	mark := make([]int, amount+1)
	sort.Slice(coins, func(i, j int) bool {return coins[i] <= coins[j]})
	var tmpMin int
	for curAmount := coins[0]; curAmount < amount+1; curAmount++ {
		tmpMin = math.MaxInt64
		for _, coin := range coins {
			if curAmount < coin {
				break
			}
			if (curAmount == coin || mark[curAmount-coin] != 0) &&
				mark[curAmount-coin] < tmpMin {
				tmpMin = mark[curAmount-coin]
			}
		}
		if tmpMin != math.MaxInt64 {
			mark[curAmount] = tmpMin + 1
		}
	}
	if mark[amount] == 0 {
		return -1
	}
	return mark[amount]
}
