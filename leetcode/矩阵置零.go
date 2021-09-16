package leetcode

func setZeroes(matrix [][]int) {
	xMark, yMark := make([]bool, len(matrix)), make([]bool, len(matrix[0]))
	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[0]); y++ {
			if matrix[x][y] == 0 {
				xMark[x] = true
				yMark[y] = true
			}
		}
	}
	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[0]); y++ {
			if xMark[x] || yMark[y] {
				matrix[x][y] = 0
			}
		}
	}
}
