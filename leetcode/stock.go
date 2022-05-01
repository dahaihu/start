package leetcode

func maxProfit1(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	preMin := prices[0]
	profit := 0
	for _, price := range prices {
		if tmp := price - preMin; tmp > profit {
			profit = tmp
		}
		if price < preMin {
			preMin = price
		}
	}
	return profit
}

// 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
func maxProfit2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	profit := 0
	pre := prices[0]
	for _, price := range prices[1:] {
		if tmp := price - pre; tmp > 0 {
			profit += tmp
		}
		pre = price
	}
	return profit
}

func maxProfit3(times int, prices []int) int {
	mark := make([][]int, times+1)
	for i := 0; i <= times; i++ {
		mark[i] = make([]int, len(prices))
	}
	for curTime := 1; curTime <= times; curTime++ {
		preMax := -prices[0]
		for idx := 1; idx < len(prices); idx++ {
			if cur := mark[curTime-1][idx-1] - prices[idx]; cur > preMax {
				preMax = cur
			}
			mark[curTime][idx] = mark[curTime][idx-1]
			if cur := preMax + prices[idx]; cur > mark[curTime][idx] {
				mark[curTime][idx] = cur
			}
		}
	}
	return mark[times][len(prices)-1]
}
