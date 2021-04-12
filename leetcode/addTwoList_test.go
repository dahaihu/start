package leetcode

import (
	"fmt"
	"testing"
)

func TestAddTwoList(t *testing.T) {
	a := initiateList([]int{2, 4, 9})
	b := initiateList([]int{5, 6, 4, 9})
	res := addTwoNumbers(a, b)
	traverseList(res)
}

func maxProfit(prices []int, times int) int {
	mark := make([][]int, times+1)
	for time := 0; time <= times; time++ {
		mark[time] = make([]int, len(prices))
	}
	for time := 1; time <= times; time++ {
		preMax := -prices[0]
		for day := 1; day < len(prices); day++ {
			if tmp := mark[time-1][day-1] - prices[day]; tmp > preMax {
				preMax = tmp
			}
			mark[time][day] = max(mark[time][day-1], prices[day]+preMax)
		}
	}
	return mark[times][len(prices)-1]
}

// [3,3,5,0,0,3,1,4]
func TestProfit(t *testing.T) {
	fmt.Println(maxProfit([]int{1,2,3,4,5}, 2))
}