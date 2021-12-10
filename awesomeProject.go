package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name string
	Age  int
}

func binarySearchLeftPosition(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (right-left)/2 + left
		if target <= nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func editDistance(word1, word2 string) int {
	m, n := len(word1), len(word2)
	mark := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		mark[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		mark[i][0] = i
	}
	for i := 1; i <= n; i++ {
		mark[0][i] = i
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				mark[i][j] = mark[i-1][j-1]
			} else {
				mark[i][j] = min(
					mark[i-1][j-1],
					min(mark[i-1][j], mark[i][j-1]),
				) + 1
			}
		}
	}
	return mark[m][n]
}

func splitArray(nums1, nums2 []int) (int, int) {
	// nums1Len < nums2Len
	nums1Len, nums2Len := len(nums1), len(nums2)
	count := (nums1Len + nums2Len + 1) / 2
	// 思考，进步，远远的超过
	left, right := 0, nums1Len
	for left < right {
		nums1Idx := (right-left)/2 + left
		nums2Idx := count - nums1Idx
		if nums1[nums1Idx] >= nums2[nums2Idx-1] {
			right = nums1Idx
		} else {
			left = nums1Idx + 1
		}
	}
	return left, count - left
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func arrayIdx(array []int, idx int, defaultVal int) int {
	if idx >= 0 && idx < len(array) {
		return array[idx]
	}
	return defaultVal
}

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	nums1Len, nums2Len := len(nums1), len(nums2)
	if nums1Len > nums2Len {
		return findMedianSortedArrays(nums2, nums1)
	}
	nums1Idx, nums2Idx := splitArray(nums1, nums2)
	leftVal := max(arrayIdx(nums1, nums1Idx-1, math.MinInt64),
		arrayIdx(nums2, nums2Idx-1, math.MinInt64))
	if (nums1Len+nums2Len)%2 == 1 {
		return float64(leftVal)
	}
	rightVal := min(arrayIdx(nums1, nums1Idx, math.MaxInt64),
		arrayIdx(nums2, nums2Idx, math.MaxInt64))
	return (float64(leftVal) + float64(rightVal)) / 2
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(splitArray(nums2, nums1))
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
