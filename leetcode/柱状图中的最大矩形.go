package leetcode

type element struct {
	idx    int
	height int
}

func largestRectangleArea(heights []int) int {
	mark := make([]*element, 0, len(heights))
	mx := 0
	for idx, h := range heights {
		if len(mark) == 0 || mark[len(mark)-1].height < h {
			mark = append(mark, &element{idx: idx, height: h})
			continue
		}
		var preIdx int
		for len(mark) != 0 && mark[len(mark)-1].height >= h {
			ele := mark[len(mark)-1]
			if tmp := (idx - ele.idx) * ele.height; tmp > mx {
				mx = tmp
			}
			preIdx = ele.idx
			mark = mark[:len(mark)-1]
		}
		mark = append(mark, &element{idx: preIdx, height: h})
	}
	for _, ele := range mark {
		if tmp := (len(heights) - ele.idx) * ele.height; tmp > mx {
			mx = tmp
		}
	}
	return mx
}

//func max(a, b int) int {
//	if a >= b {
//		return a
//	}
//	return b
//}
