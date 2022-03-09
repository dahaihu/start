package leetcode

type element struct {
	idx    int
	height int
}

func largestRectangleArea(heights []int) int {
	var (
		mark   []*element
		result int
	)
	for idx, height := range heights {
		preIdx := idx
		for len(mark) > 0 && mark[len(mark)-1].height > height {
			preItem := mark[len(mark)-1]
			if area := (idx - preItem.idx) * preItem.height; area > result {
				result = area
			}
			preIdx = preItem.idx
			mark = mark[:len(mark)-1]
		}
		mark = append(mark, &element{idx: preIdx, height: height})
	}
	for _, ele := range mark {
		if area := (len(heights) - ele.idx) * ele.height; area > result {
			result = area
		}
	}
	return result
}

//func max(a, b int) int {
//	if a >= b {
//		return a
//	}
//	return b
//}
