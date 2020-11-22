package leetcode

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	fillData(0, 0, n, n, 1, res)
	return res
}

func fillData(x, y, lengthX, lengthY, cur int, arr [][]int) {
	// 矩形上边
	for i := y; i <= y+lengthY-1; i++ {
		arr[x][i] = cur
		cur += 1
	}
	if lengthX == 1 {
		return
	}
	// 矩形右边
	for i := x + 1; i < x+lengthX-1; i++ {
		arr[i][y+lengthY-1] = cur
		cur += 1
	}
	// 矩形下边
	for i := y + lengthY - 1; i >= y; i-- {
		arr[x+lengthX-1][i] = cur
		cur += 1
	}
	if lengthY == 2 {
		return
	}
	// 矩阵左边
	for i := x + lengthX - 2; i >= x+1; i-- {
		arr[i][y] = cur
		cur += 1
	}
	fillData(x+1, y+1, lengthX-2, lengthY-2, cur, arr)
}
