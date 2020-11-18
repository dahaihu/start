package leetcode

func solveNQueens(n int) [][]int {
	mark := make([][]int, 0)
	var solveQueensIdx func([]int, int)
	solveQueensIdx = func(tmpRes []int, idx int) {
		//fmt.Println("tmpRes is ", tmpRes, "idx is ", idx)
		if idx == n {
			mark = append(mark, tmpRes)
			return
		}
		for i := 0; i < n; i++ {
			success := true
			for prev, num := range tmpRes {
				if i == num || idx - prev == abs(i-num) {
					success = false
					break
				}
			}
			if success {
				tmp := make([]int, len(tmpRes) + 1)
				copy(tmp[:len(tmpRes)], tmpRes)
				tmp[len(tmpRes)] = i
				solveQueensIdx(tmp, idx+1)
			}
		}
	}
	solveQueensIdx([]int{}, 0)
	return mark
}

func abs(val int) int {
	if val >= 0 {
		return val
	}
	return -val
}