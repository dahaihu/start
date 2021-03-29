package leetcode


func change(amount int, coins []int) int {
	mark := make([]int, amount+1)
	mark[0] = 1
	for _, coin := range coins {
		for curAmount := coin; curAmount <= amount; curAmount++ {
			mark[curAmount] = mark[curAmount] + mark[curAmount-coin]
		}
	}
	return mark[amount]
}