package leetcode

/**
* @Author: 胡大海
* @Date: 2020-03-23 08:53
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

// 感觉堆排序，特别有意思，不知道为啥，就是觉得有意思

// 时间复杂度应该如何取证呢？？？

func adjust(idx int, margin int, array []int) {
	maxIdx := 2*idx + 1
	if maxIdx >= margin {
		return
	}
	if tmp := maxIdx + 1; tmp < margin && array[tmp] > array[maxIdx] {
		maxIdx = tmp
	}
	if array[idx] >= array[maxIdx] {
		return
	}

	tmp := array[maxIdx]
	array[maxIdx] = array[idx]
	array[idx] = tmp

	adjust(maxIdx, margin, array)

}

func HeapSort(array []int) {
	length := len(array)
	for idx := (length - 1) / 2; idx >= 0; idx-- {
		adjust(idx, length, array)
	}
	for idx := length - 1; idx >= 0; idx-- {
		tmp := array[0]
		array[0] = array[idx]
		array[idx] = tmp
		adjust(0, idx, array[:idx])
	}

}
