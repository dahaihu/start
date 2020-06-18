package leetcode

/**
* @Author: 胡大海
* @Date: 2019-10-29 08:12
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

func singleNumber(nums []int) []int {
	mark := 0
	for _, num := range nums {
		mark ^= num
	}
	// 感觉这个操作好骚，要是我只会向右一个一个移动，找到第一个非0的值
	mark = mark & ((^mark) + 1)
	nums1 := 0
	nums2 := 0
	for _, num := range nums {
		if mark&num == 0 {
			nums1 ^= num
		} else {
			nums2 ^= num
		}
	}
	return []int{nums1, nums2}
}
