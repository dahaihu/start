package leetcode

type element struct {
	idx    int
	height int
}

func largestRectangleArea(heights []int) int {

	mark := make([]element, 0)
	mx := 0
	for idx, height := range heights {
		if len(mark) == 0 {
			mark = append(mark, element{idx: idx, height: height})
			continue
		}
		for ele := mark[len(mark)-1]; ele.height <= height && len(mark) != 0; mark = mark[:len(mark)-1] {
			if sum := (idx - ele.idx) * ele.height; sum > mx {
				mx = sum
			}
		}
		mark = append(mark, element{idx: idx, height: height})
	}

	for _, ele := range mark {
		if sum := ele.height * (len(heights) - ele.idx); sum > mx {
			mx = sum
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
