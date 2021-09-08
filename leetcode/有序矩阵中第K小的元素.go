package leetcode

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	left, right := matrix[0][0], matrix[n-1][n-1]
	for left < right {
		mid := (right-left)/2 + left
		if num := leNum(matrix, mid, n); num < k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}

func leNum(matrix [][]int, target, n int) int {
	var num int
	xIdx, yIdx := n-1, 0
	for xIdx >= 0 && yIdx < n {
		if target >= matrix[xIdx][yIdx] {
			yIdx += 1
			num += xIdx + 1
		} else {
			xIdx -= 1
		}
	}
	return num
}
