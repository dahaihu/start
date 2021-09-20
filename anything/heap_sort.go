package anything

/**
* @Author: 胡大海
* @Date: 2020-10-03 15:16
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

func adjust(nums []int, start, end int) {
	var maxChild int
	if maxChild = 2*start + 1; maxChild >= end {
		return
	}
	if rightChild := maxChild + 1;
		rightChild < end &&
			nums[rightChild] > nums[maxChild] {
		maxChild = rightChild
	}
	if nums[start] >= nums[maxChild] {
		return
	}
	nums[start], nums[maxChild] = nums[maxChild], nums[start]
	adjust(nums, maxChild, end)
}
func heapSort(nums []int) {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		adjust(nums, i, len(nums))
	}
	for i := len(nums)-1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		adjust(nums, 0, i)
	}
}
