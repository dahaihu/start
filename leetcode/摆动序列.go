package leetcode

func status(val int) int {
	if val > 0 {
		return 1
	} else if val < 0 {
		return -1
	} else {
		return 0
	}
}

func wiggleMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 1
	preStatus := 100
	for i := 1; i < len(nums); i++ {
		curStatus := status(nums[i] - nums[i-1])
		if curStatus == preStatus || curStatus == 0 {
			continue
		}
		preStatus = curStatus
		count += 1
	}
	return count
}
