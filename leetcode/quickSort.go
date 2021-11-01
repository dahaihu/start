package leetcode

func sub(nums []int, start, end int) int {
	left, right := start+1, end
	for {
		for left <= right && nums[left] <= nums[start] {
			left++
		}
		for left <= right && nums[right] >= nums[start] {
			right--
		}
		if left > right {
			break
		}
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	nums[start], nums[right] = nums[right], nums[start]
	return right
}

func quickSort(nums []int, start, end int) {
	if start < end {
		mid := sub(nums, start, end)
		quickSort(nums, start, mid-1)
		quickSort(nums, mid+1, end)
	}
}

func quickMain(nums []int) {
	quickSort(nums, 0, len(nums)-1)
	//fmt.Println(nums)
}
