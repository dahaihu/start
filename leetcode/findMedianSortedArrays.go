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
