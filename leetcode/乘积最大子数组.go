package leetcode

func maxProduct(nums []int) int {
	preMin, preMax, mx := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		preMin, preMax =
			arrayMin([]int{preMin * nums[i], preMax * nums[i], nums[i]}),
			arrayMax([]int{preMin * nums[i], preMax * nums[i], nums[i]})
		if preMax > mx {
			mx = preMax
		}
	}
	return mx
}

func arrayMax(nums []int) int {
	mx := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > mx {
			mx = nums[i]
		}
	}
	return mx
}

func arrayMin(nums []int) int {
	mn := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < mn {
			mn = nums[i]
		}
	}
	return mn
}
