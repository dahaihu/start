package leetcode

/**
* @Author: 胡大海
* @Date: 2020-03-23 08:53
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

// 感觉堆排序，特别有意思，不知道为啥，就是觉得有意思

// 时间复杂度应该如何取证呢？？？

func adjust(idx int, margin int, array []int) {
	left := 2*idx + 1
	if left >= margin {
		return
	}
	mxChild := left
	if right := left + 1; right < margin && array[right] > array[left] {
		mxChild = right
	}
	if array[idx] >= array[mxChild] {
		return
	}
	array[idx], array[mxChild] = array[mxChild], array[idx]
	adjust(mxChild, margin, array)

}

func HeapSort(array []int) {
	for i := len(array)/2 - 1; i >= 0; i-- {
		adjust(i, len(array), array)
	}
	for i := len(array) - 1; i >= 0; i-- {
		array[0], array[i] = array[i], array[0]
		adjust(0, i, array)
	}
}
