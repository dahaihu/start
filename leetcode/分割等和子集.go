package leetcode

import "fmt"

func canPartition(nums []int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	dp := make([][]bool, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		dp[i] = make([]bool, target+1)
	}
	dp[0][0] = true
	for ni := 1; ni <= len(nums); ni++ {
		for ti := 0; ti <= target; ti++ {
			dp[ni][ti] = dp[ni-1][ti]
			if nums[ni-1] == ti {
				dp[ni][ti] = true
			} else if nums[ni-1] < ti {
				dp[ni][ti] = dp[ni][ti] || dp[ni-1][ti-nums[ni-1]]
			}
		}
	}
	fmt.Println(dp)
	return dp[len(nums)][target]
}
