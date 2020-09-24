package leetcode

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

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func dividedTwoSortedArrays(nums1, nums2 []int) (int, int) {
	m, n := len(nums1), len(nums2)
	midAll := (m + n + 1) / 2
	left, right := 0, m
	for left < right {
		p := (right-left)/2 + left
		k := midAll - p
		if nums2[k-1] < nums1[p] {
			right = p
		} else {
			left = p + 1
		}
	}
	return left, midAll - left
}

// k 表示一个正整数, end是包含在边界内的数据
func findKth(nums1 []int, start1, end1 int, nums2 []int, start2, end2 int, k int) int {
	if end1-start1 > end2-start2 {
		return findKth(nums2, start2, end2, nums1, start1, end1, k)
	}
	if start1 > end1 {
		return nums2[start2+k-1]
	}
	if k == 1 {
		return min(nums1[start1], nums2[start2])
	}
	nums1Mid := min(start1+(k/2)-1, end1)
	nums2Mid := min(start2+(k/2)-1, end2)
	if nums1[nums1Mid] <= nums2[nums2Mid] {
		return findKth(nums1, nums1Mid+1, end1, nums2, start2, end2, k-(nums1Mid-start1+1))
	}
	return findKth(nums1, start1, end1, nums2, nums2Mid+1, end2, k-(nums2Mid-start2+1))
}

func FindMedianSortedArrays(nums1, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if (m+n)%2 == 1 {
		return float64(findKth(nums1, 0, len(nums1)-1, nums2, 0, len(nums2)-1, (m+n)/2+1))
	}
	return float64(findKth(nums1, 0, len(nums1)-1, nums2, 0, len(nums2)-1, (m+n)/2)+findKth(nums1, 0, len(nums1)-1, nums2, 0, len(nums2)-1, (m+n)/2+1)) / 2

}
