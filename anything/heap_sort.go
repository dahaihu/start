package anything

/**
* @Author: 胡大海
* @Date: 2020-10-03 15:16
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

func adjust(nums []int, idx, end int) {
	leftIdx := 2*idx + 1
	if leftIdx >= end {
		return
	}
	maxIdx := leftIdx
	if rightIdx := leftIdx + 1; rightIdx < end && nums[rightIdx] > nums[maxIdx] {
		maxIdx = rightIdx
	}

	if nums[idx] >= nums[maxIdx] {
		return
	}
	nums[idx], nums[maxIdx] = nums[maxIdx], nums[idx]
	adjust(nums, maxIdx, end)
}

func heapSort(nums []int) {
	length := len(nums)
	for i := length/2 - 1; i >= 0; i-- {
		adjust(nums, i, length)
	}
	for i := length - 1; i >= 1; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		adjust(nums, 0, i)
	}
}
