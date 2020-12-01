package leetcode

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	fillData(0, 0, n, 1, matrix)
	return matrix
}

func fillData(x, y, length, cur int, matrix [][]int) {
	for i := y; i < y+length; i++ {
		matrix[x][i] = cur
		cur += 1
	}
	if length == 1 {
		return
	}
	for i := x + 1; i < x+length-1; i++ {
		matrix[i][y+length-1] = cur
		cur += 1
	}

	for i := y + length - 1; i >= y; i-- {
		matrix[x+length-1][i] = cur
		cur += 1
	}

	if length == 2 {
		return
	}

	for i := x + length - 2; i > x; i-- {
		matrix[i][y] = cur
		cur += 1
	}

	fillData(x+1, y+1, length-2, cur, matrix)
}
