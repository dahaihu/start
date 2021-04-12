package leetcode

import "fmt"

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	maxLength := 0
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
		for j, ele := range matrix[i] {
			dp[i][j] = int(ele - '0')
			if dp[i][j] == 1 {
				maxLength = 1
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
				if dp[i][j] > maxLength {
					maxLength = dp[i][j]
				}
			}
		}
	}
	fmt.Println("dp is ")
	fmt.Println(dp)
	return maxLength * maxLength
}
