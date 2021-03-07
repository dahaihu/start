package leetcode

func solveNQueens(n int) [][]int {
	res := [][]int{}
	var dfs func([]int, int)
	dfs = func(internalRes []int, idx int) {
		if idx == n {
			res = append(res, internalRes)
			return
		}
		for val := 0; val < n; val++ {
			if conflicts(internalRes, idx, val) {
				continue
			}
			tmp := make([]int, len(internalRes)+1)
			copy(tmp, internalRes)
			tmp[len(internalRes)] = val
			dfs(tmp, idx+1)
		}
	}
	dfs([]int{}, 0)
	return res
}

func conflicts(prev []int, idx, val int) bool {
	for prevIdx, prevVal := range prev {
		if val == prevVal || (idx-prevIdx) == abs(prevVal - val) {
			return true
		}
	}
	return false
}

func abs(val int) int {
	if val >= 0 {
		return val
	}
	return -val
}
