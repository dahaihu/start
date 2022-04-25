package leetcode

func canPartition(nums []int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	mark := make([][]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		mark[i] = make([]bool, target)
	}
	if nums[0] <= target {
		mark[0][nums[0]-1] = true
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < target; j++ {
			mark[i][j] = mark[i-1][j]
			if nums[i] == j+1 {
				mark[i][j] = true
			} else if nums[i] < j+1 {
				mark[i][j] = mark[i][j] || mark[i-1][j+1-nums[i]-1]
			}
		}
	}
	return mark[len(nums)-1][target-1]
}
