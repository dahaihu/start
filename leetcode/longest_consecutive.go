package leetcode

import "fmt"

func longestConsecutive(nums []int) int {
	mark := make(map[int]int)
	res := 0
	for _, num := range nums {
		if _, ok := mark[num]; !ok {
			left := mark[num-1]
			right := mark[num+1]
			cur := left + right + 1
			if cur > res {
				res = cur
			}
			// 这个地方的更新还是很有意思的？
			// 什么场景进行更新呢？
			//mark[num] = cur
			mark[num] = cur
			mark[num-left] = cur
			mark[num+right] = cur
			fmt.Println("updated mark", mark)
		}
	}
	return res
}