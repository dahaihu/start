package leetcode

import (
	"math"
	"sort"
)

// how to initiate
func coinChange(coins []int, amount int) int {
	mark := make([]int, amount+1)
	for idx := 1; idx <= amount; idx++ {
		mark[idx] = -1
	}
	sort.Slice(coins, func(i, j int) bool {return coins[i] <= coins[j]})
	for curAmount := coins[0]; curAmount <= amount; curAmount++ {
		minCnt := math.MaxInt64
		for _, coin :=range coins {
			if coin > curAmount {
				break
			}
			if tmp := mark[curAmount-coin]; tmp != -1 && tmp < minCnt {
				minCnt = tmp
			}
		}
		if minCnt != math.MaxInt64 {
			mark[curAmount] = minCnt + 1
		}
	}
	return mark[amount]
}

func coinChangeFromTop(coins []int, amount int) int {
	const NOTPROCESSED = -2
	const NOTREACHABLE = -1
	mark := make([]int, amount+1)
	for idx := 1; idx <= amount; idx++ {
		mark[idx] = NOTPROCESSED
	}
	sort.Slice(coins, func(i, j int) bool { return coins[i] <= coins[j] })
	var dp func(int) int
	dp = func(curAmount int) int {
		if mark[curAmount] == NOTREACHABLE {
			return math.MaxInt64
		} else if mark[curAmount] != NOTPROCESSED {
			return mark[curAmount]
		}
		min := math.MaxInt64
		for _, coin := range coins {
			if coin > curAmount {
				break
			}
			if tmp := dp(curAmount - coin); tmp < min && tmp >= 0 {
				min = tmp
			}
		}
		if min == math.MaxInt64 {
			mark[curAmount] = NOTREACHABLE
		} else {
			mark[curAmount] = min + 1
		}
		return mark[curAmount]
	}
	tmp := dp(amount)
	return tmp
}

func coinChangeUsingDfs(coins []int, amount int) int {
	sort.Slice(coins, func(i, j int) bool {return coins[i]>=coins[j]})
	coinCnt := math.MaxInt64
	var dep func([]int, int, int, *int)
	dep = func(coins []int, count int, amount int, coinCnt *int) {
		if amount == 0 {
			if count < *coinCnt {
				*coinCnt = count
			}
			return
		}
		if len(coins) == 0 {
			return
		}
		for cnt := amount/coins[0]; cnt >= 0 && cnt + count < *coinCnt; cnt-- {
			dep(coins[1:], count + cnt, amount - cnt*coins[0], coinCnt)
		}

	}
	dep(coins, 0, amount, &coinCnt)
	if coinCnt == math.MaxInt64 {
		return -1
	}
	return coinCnt
}
