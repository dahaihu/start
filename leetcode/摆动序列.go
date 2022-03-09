package leetcode

func wiggleMaxLength(nums []int) int {
	const (
		bigger  = 1
		equal   = 0
		smaller = -1
	)
	preStatus := 100
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		var status int
		if nums[i] > nums[i-1] {
			status = bigger
		} else {
			status = smaller
		}
		if preStatus == status {
			continue
		} else {
			count += 1
			preStatus = status
		}
	}
	return count
}
