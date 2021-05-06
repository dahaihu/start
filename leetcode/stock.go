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
		if tmp := price-pre; tmp > 0 {
			profit += tmp
		}
		pre = price
	}
	return profit
}

func maxProfit3(times int, prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	mark := make([][]int, times+1)
	for i := 0; i <= times; i++ {
		mark[i] = make([]int, len(prices))
	}
	for time := 1; time <= times; time++ {
		preMax := -prices[0]
		for day := 1; day < len(prices); day++ {
			if tmp := mark[time-1][day-1] - prices[day]; tmp > preMax {
				preMax = tmp
			}
			if tmp := preMax + prices[day]; tmp > mark[time][day-1] {
				mark[time][day] = tmp
			} else {
				mark[time][day] = mark[time][day-1]
			}
		}
	}
	return mark[times][len(prices)-1]
}
