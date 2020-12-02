package leetcode

import (
	"fmt"
	"math"
	"testing"
)

func splitArrays(nums1, nums2 []int) (int, int) {
	m, n := len(nums1), len(nums2)
	middleCount := (m + n) / 2
	left, right := 0, m
	for left < right {
		mid1 := (right-left)/2 + left
		if nums1[mid1] >= nums2[middleCount-mid1-1] {
			right = mid1
		} else {
			left = mid1 + 1
		}
	}
	return right, middleCount - right
}

func findMedianSortedArrayUsingSplit(nums1, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrayUsingSplit(nums2, nums1)
	}
	m, n := len(nums1), len(nums2)

	mid1, mid2 := splitArrays(nums1, nums2)

	nums1LeftMargin := getValueUsingIndex(nums1, mid1-1, math.MinInt64)
	nums2LeftMargin := getValueUsingIndex(nums2, mid2-1, math.MinInt64)

	nums1RightMargin := getValueUsingIndex(nums1, mid1, math.MaxInt64)
	nums2RightMargin := getValueUsingIndex(nums2, mid2, math.MaxInt64)
	if (m+n)%2 == 0 {
		return float64(max(nums1LeftMargin, nums2LeftMargin) + min(nums1RightMargin, nums2RightMargin))/2
	}
	return float64(min(nums2RightMargin, nums1RightMargin))
}

func getValueUsingIndex(nums []int, index, defaultValue int) int {
	if index >= 0 && index < len(nums) {
		return nums[index]
	}
	return defaultValue
}

func TestGenerateMatrix(t *testing.T) {
	fmt.Println(generateMatrix(3))
}

func TestQQuickSort(t *testing.T) {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{2, 3, 5, 6, 6}
	fmt.Println(findMedianSortedArrayUsingSplit(nums2, nums1))
}
