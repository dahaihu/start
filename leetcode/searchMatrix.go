package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	array := make([]int, m)
	for i := 0; i < m; i++ {
		array[i] = matrix[i][n-1]
	}
	row := universalBinarySearch(array, target)
	if row == m {
		return false
	}
	column := universalBinarySearch(matrix[row], target)
	if column >= 0 && column < n && matrix[row][column] == target {
		return true
	}
	return false
}

func universalBinarySearch(array []int, target int) int {
	left, right := 0, len(array)
	for left < right {
		mid := (right-left)/2 + left
		if target <= array[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
