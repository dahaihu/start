package leetcode

import (
	"strconv"
	"strings"
)

type Height struct {
	index  int
	height int
}

func trap(height []int) int {
	waters := 0
	mark := make([]Height, 0)
	for idx, h := range height {
		preHeight := 0
		for len(mark) != 0 && mark[len(mark)-1].height < h {
			lastItem := mark[len(mark)-1]
			waters += (idx - lastItem.index - 1) * (lastItem.height - preHeight)
			preHeight = lastItem.height
			mark = mark[:len(mark)-1]
		}
		if len(mark) == 0 {
			mark = append(mark, Height{index: idx, height: h})
			continue
		}
		lastItem := mark[len(mark)-1]
		if lastItem.height == h {
			waters += (idx - lastItem.index - 1) * (lastItem.height - preHeight)
			mark[len(mark)-1].index = idx
		} else {
			waters += (idx - lastItem.index - 1) * (h - preHeight)
			mark = append(mark, Height{index: idx, height: h})
		}
	}
	return waters
}

func getPermutation(n int, k int) string {
	mark := make([]int, n)
	mark[0] = 1
	for i := 1; i < n; i++ {
		mark[i] = mark[i-1] * i
	}

	items := make([]string, n, n)
	for i := 0; i < n; i++ {
		items[i] = strconv.Itoa(i + 1)
	}

	res := make([]string, 0, n)
	var idx int
	for {
		idx, k = k/mark[n-1], k%mark[n-1]
		if k == 0 {
			res = append(res, items[idx-1])
			items = popKItem(idx-1, items)
			for i := len(items) - 1; i >= 0; i-- {
				res = append(res, items[i])
			}
			break
		}
		res = append(res, items[idx])
		items = popKItem(idx, items)
		n = n - 1

	}
	return strings.Join(res, "")
}

func popKItem(idx int, items []string) []string {
	res := make([]string, len(items)-1)
	copy(res[:idx], items[:idx])
	copy(res[idx:], items[idx+1:])
	return res
}
