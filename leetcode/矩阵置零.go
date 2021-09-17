package leetcode

import "fmt"

func setZeroes(matrix [][]int) {
	var row0ExistZero, column0ExistZero bool
	for column := 0; column < len(matrix[0]); column++ {
		if matrix[0][column] == 0 {
			row0ExistZero = true
			break
		}
	}
	for row := 0; row < len(matrix); row++ {
		if matrix[row][0] == 0 {
			column0ExistZero = true
			break
		}
	}

	for row := 0; row < len(matrix); row++ {
		for column := 0; column < len(matrix[0]); column++ {
			if matrix[row][column] == 0 {
				matrix[row][0] = 0
				matrix[0][column] = 0
			}
		}
	}
	fmt.Println("first step result is ", matrix)
	for row := 1; row < len(matrix); row++ {
		for column := 1; column < len(matrix[0]); column++ {
			if matrix[row][0] == 0 || matrix[0][column] == 0 {
				matrix[row][column] = 0
			}
		}
	}
	fmt.Println("second step result is ", matrix)
	if row0ExistZero {
		for column := 0; column < len(matrix[0]); column++ {
			matrix[0][column] = 0
		}
	}

	if column0ExistZero {
		for row := 0; row < len(matrix); row++ {
			matrix[row][0] = 0
		}
	}

}
