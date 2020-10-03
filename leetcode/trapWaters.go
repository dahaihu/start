package leetcode

type Height struct {
	index  int
	height int
}

func trap(height []int) int {
	waters := 0
	mark := make([]Height, 0)
	for idx, h := range height {
		preHeight := 0
		for len(mark) != 0 {
			i := len(mark) - 1
			if mark[i].height > h {
				waters += (h - preHeight) * (idx - mark[i].index - 1)
				mark = append(mark, Height{idx, h})
				break
			} else if mark[i].height == h {
				waters += (h - preHeight) * (idx - mark[i].index - 1)
				mark[i].index = idx
				break
			} else {
				waters += (mark[i].height - preHeight) * (idx - mark[i].index - 1)
				preHeight = mark[i].height
				mark = mark[:i]
			}
		}
		if len(mark) == 0 {
			mark = append(mark, Height{idx, h})
		}
	}
	return waters
}
