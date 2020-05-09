package leetcode

import (
	"bytes"
	"strconv"
)

func getPermutation(n int, k int) string {
	mark := make([]int, n)
	mark[0] = 1
	for i := 1; i < n; i++ {
		mark[i] = i * mark[i-1]
	}
	items := make([]int, n)
	for i := 0; i < n; i++ {
		items[i] = i + 1
	}
	var idx, item int
	res := bytes.Buffer{}

	for true {
		idx, k = k/mark[len(items)-1], k%mark[len(items)-1]
		if k == 0 {
			items, item = popKItem(items, idx-1)
			res.WriteString(strconv.Itoa(item))
			for i := len(items) - 1; i >= 0; i-- {
				res.WriteString(strconv.Itoa(items[i]))
			}
			break
		}
		items, item = popKItem(items, idx)
		res.WriteString(strconv.Itoa(item))
		if k == 1 {
			for i := 0; i <= len(items) - 1; i++ {
				res.WriteString(strconv.Itoa(items[i]))
			}
			break
		}
	}


	return res.String()
}
func popKItem(items []int, idx int) ([]int, int) {
	item := items[idx]
	newItems := make([]int, idx)
	copy(newItems, items[:idx])
	newItems = append(newItems, items[idx+1:]...)
	return newItems, item
}
