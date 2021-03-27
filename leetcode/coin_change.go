package leetcode

import (
	"math"
	"sort"
)

// how to initiate
func coinChange(coins []int, amount int) int {
	mark := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		mark[i] = -1
	}
	sort.Slice(coins, func(i, j int) bool { return coins[i] <= coins[j] })
	for curAmount := coins[0]; curAmount <= amount; curAmount++ {
		prevMinUsedCnt := math.MaxInt64
		for _, coin := range coins {
			if coin > curAmount {
				break
			}
			if tmp := mark[curAmount-coin]; tmp >= 0 && tmp < prevMinUsedCnt {
				prevMinUsedCnt = tmp
			}
		}
		if prevMinUsedCnt != math.MaxInt64 {
			mark[curAmount] = prevMinUsedCnt + 1
		}
	}
	return mark[amount]
}

func coinChangeFromTop(coins []int, amount int) int {
	mark := make([]int, amount+1)
	const NotProcessed = -2
	const NotReachable = -1
	for i := 1; i <= amount; i++ {
		mark[i] = NotProcessed
	}
	sort.Slice(coins, func(i, j int) bool { return coins[i] <= coins[j] })
	var dp func(int) int
	dp = func(curAmount int) int {
		if mark[curAmount] != NotProcessed {
			return mark[curAmount]
		}
		prevMinUsedCnt := math.MaxInt64
		for _, coin := range coins {
			if coin > curAmount {
				break
			}
			if tmp := dp(curAmount - coin); tmp < prevMinUsedCnt && tmp >= 0 {
				prevMinUsedCnt = tmp
			}
		}
		if prevMinUsedCnt == math.MaxInt64 {
			mark[curAmount] = NotReachable
		} else {
			mark[curAmount] = prevMinUsedCnt + 1
		}
		return mark[curAmount]
	}
	return dp(amount)
}

func coinChangeUsingDfs(coins []int, amount int) int {
	sort.Slice(coins, func(i, j int) bool { return coins[i] >= coins[j] })
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
		for cnt := amount / coins[0]; cnt >= 0 && cnt+count < *coinCnt; cnt-- {
			dep(coins[1:], count+cnt, amount-cnt*coins[0], coinCnt)
		}

	}
	dep(coins, 0, amount, &coinCnt)
	if coinCnt == math.MaxInt64 {
		return -1
	}
	return coinCnt
}
