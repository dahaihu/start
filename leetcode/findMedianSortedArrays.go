package leetcode

import (
	"math"
)

/**
* @Author: 胡大海
* @Date: 2019-11-03 11:06
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

/*
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

// 傻逼解法并没有满足要要求，人家的要求是O(log(m + n)),而我这个解法是O(m+n)，哎
// go中的除法很有意思， 3 / 2 = 1, 3.0 / 2 = 1.5，这个是做算法题需要考虑的东西
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mark := make([]int, 0)
	for ; len(nums1) != 0 && len(nums2) != 0; {
		if nums1[0] <= nums2[0] {
			mark = append(mark, nums1[0])
			nums1 = nums1[1:]
		} else {
			mark = append(mark, nums2[0])
			nums2 = nums2[1:]
		}
	}

	if len(nums1) != 0 {
		mark = append(mark, nums1...)
	} else {
		mark = append(mark, nums2...)
	}

	length := len(mark)
	if length&1 == 1 {
		return float64(mark[(length-1)/2])
	} else {
		return float64(mark[length/2]+mark[(length/2)-1]) / 2
	}
}

func FindMedianSortedArraysAnswer(nums1, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	idx, remain := (m+n-1)/2, (m+n-1)%2
	if remain == 0 {
		return float64(findKth(nums1, nums2, 0, m-1, 0, n-1, idx))
	}
	return float64(findKth(nums1, nums2, 0, m-1, 0, n-1, idx)+findKth(nums1, nums2, 0, m-1, 0, n-1, idx+1)) / 2
}

func findKth(nums1, nums2 []int, start1, end1, start2, end2, k int) int {
	if end1-start1 > end2-start2 {
		return findKth(nums2, nums1, start2, end2, start1, end1, k)
	}
	if start1 > end1 {
		return nums2[start2+k]
	}
	if k == 0 {
		return minAB(nums1[start1], nums2[start2])
	}
	mid1 := minAB(end1, start1+((k+1)/2)-1)
	mid2 := minAB(end2, start2+((k+1)/2)-1)
	if nums1[mid1] > nums2[mid2] {
		return findKth(nums1, nums2, start1, end1, mid2+1, end2, k-(mid2-start2+1))
	}
	return findKth(nums1, nums2, mid1+1, end1, start2, end2, k-(mid1-start1+1))
}

func minAB(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func DivideSortedArrays(nums1, nums2 []int) (int, int) {
	m, n := len(nums1), len(nums2)
	k := (m + n + 1) / 2
	left, right := 0, m
	for left < right {
		mid1 := (right - left) / 2 + left
		mid2 := k - mid1 - 1
		if nums1[mid1] < nums2[mid2] {
			if left == mid1 {
				return left, k - left - 2
			}
			left = mid1
		} else {
			right = mid1
		}
	}
	return right-1, k - right - 1
}

func getIdx(idx int, nums1, nums2 []int) int {
	if idx + 1 <= len(nums1) {
		return nums1[idx]
	}
	return nums2[idx - len(nums1)]
}

func GetMiddle(nums1, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	mid := (m + n + 1) / 2
	if (m + n + 1) % 2 == 0 {
		return float64(getIdx(mid, nums1, nums2))
	}
	return float64(getIdx(mid, nums1, nums2) + getIdx(mid+1, nums1, nums2)) / 2

}

func getIdxDefault(nums []int, idx int, _default int) int {
	if idx >= 0 && idx+1 <= len(nums) {
		return nums[idx]
	}
	return _default
}

func Divided(nums1, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		return Divided(nums2, nums1)
	}
	//if nums1[0] >= nums2[n-1] {
	//	return GetMiddle(nums2, nums1)
	//}
	//if nums2[0] >= nums1[m-1] {
	//	return GetMiddle(nums1, nums2)
	//}
	nums1Idx, nums2Idx := DivideSortedArrays(nums1, nums2)
	if (m + n + 1) % 2 == 0 {
		return float64(max(getIdxDefault(nums1, nums1Idx, math.MinInt64), getIdxDefault(nums2, nums2Idx, math.MinInt64)))
	}
	return float64(max(getIdxDefault(nums1, nums1Idx, math.MinInt64), getIdxDefault(nums2, nums2Idx, math.MinInt64)) + minAB(getIdxDefault(nums1, nums1Idx+1, math.MaxInt64), getIdxDefault(nums2, nums2Idx+1, math.MaxInt64))) / 2
}

func FindMedianSortedArrays(nums1, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		return FindMedianSortedArrays(nums2, nums1)
	}
	k := (m + n + 1) / 2
	left, right := 0, m
	for left < right {
		nums1Mid := left + (right-left)/2
		nums2Mid := k - nums1Mid
		if nums1[nums1Mid] < nums2[nums2Mid-1] {
			left = nums1Mid + 1
		} else {
			right = nums1Mid
		}
	}
	nums1Mid, nums2Mid := left, k-left
	nums1Margin := math.MinInt64
	nums2Margin := math.MinInt64
	if nums1Mid > 0 {
		nums1Margin = nums1[nums1Mid-1]
	}
	if nums2Mid > 0 {
		nums2Margin = nums2[nums2Mid-1]
	}
	var c1 int
	if nums1Margin > nums2Margin {
		c1 = nums1Margin
	} else {
		c1 = nums2Margin
	}
	if (m+n+1)%2 == 0 {
		return float64(c1)
	}

	nums1Margin = math.MaxInt64
	nums2Margin = math.MaxInt64

	if nums1Mid < m {
		nums1Margin = nums1[nums1Mid]
	}

	if nums2Mid < n {
		nums2Margin = nums2[nums2Mid]
	}

	var c2 int
	if nums1Margin > nums2Margin {
		c2 = nums2Margin
	} else {
		c2 = nums1Margin
	}
	return float64(c1+c2) / 2

}

func getIdxValue(nums []int, idx int, defaultValue int) int {
	if idx >= 0 && idx < len(nums) {
		return nums[idx]
	}
	return defaultValue
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
