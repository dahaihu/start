package leetcode

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	left, right := matrix[0][0], matrix[n-1][n-1]
	for left < right {
		mid := (right-left)/2 + left
		if check(matrix, mid, k, n) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func check(matrix [][]int, mid, k, n int) bool {
	// 从左下角到右上角查找
	i, j := n-1, 0
	num := 0
	for i >= 0 && j <= n-1 {
		if matrix[i][j] <= mid {
			j += 1
			num += i+1
		} else {
			i -= 1
		}
	}
	return num < k
}
