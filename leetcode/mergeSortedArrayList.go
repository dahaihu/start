package leetcode

/**
* @Author: 胡大海
* @Date: 2020-11-13 17:22
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

func shiftRight(nums []int, idx int) {
	for end := len(nums)-1; end > idx; end-- {
		nums[end] = nums[end-1]
	}
}

func mergeSortedArrayList(arr1, arr2 []int) {
	idx1, idx2 := 0, 0
	for idx2 < len(arr2) {
		if len(arr2)-idx2 == len(arr1)-idx1 {
			copy(arr1[idx1:], arr2[idx2:])
			return
		}
		if arr1[idx1] > arr2[idx2] {
			shiftRight(arr1, idx1)
			arr1[idx1] = arr2[idx2]
			idx1 += 1
			idx2 += 1
		} else {
			idx1 += 1
		}
	}
}
