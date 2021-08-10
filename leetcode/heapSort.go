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

func heapSort(array []int) {
	for i := len(array)/2 - 1; i >= 0; i-- {
		adjust(i, len(array), array)
	}
	for i := len(array) - 1; i > 0; i-- {
		array[0], array[i] = array[i], array[0]
		adjust(0, i, array)
	}
}

func selectHeapSort(nums []int) {
	// first step: make nums to big heap
	for i := 1; i < len(nums); i++ {
		cur, num := i, nums[i]
		for cur != 0 {
			parent := (cur - 1) / 2
			if nums[parent] >= num {
				break
			}
			nums[cur] = nums[parent]
			cur = parent
		}
		nums[cur] = num
	}
	for end := len(nums) - 1; end >= 0; end-- {
		nums[end], nums[0] = nums[0], nums[end]
		cur, num := 0, nums[0]
		for cur < end {
			maxChild := 2*cur + 1
			if maxChild >= end {
				break
			}
			if rightChild := maxChild + 1;
				rightChild < end && nums[rightChild] > nums[maxChild] {
				maxChild = rightChild
			}
			if nums[maxChild] <= num {
				break
			}
			cur, nums[cur] = maxChild, nums[maxChild]
		}
		nums[cur] = num
	}
}
