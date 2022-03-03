package leetcode

const (
	red   = 0
	white = 1
	blue  = 2
)

func sortColors(nums []int) {
	redPos, bluePos := 0, len(nums)-1
	var idx int
	for idx <= bluePos {
		if nums[idx] == red {
			nums[idx], nums[redPos] = nums[redPos], nums[idx]
			idx += 1
			redPos += 1
		} else if nums[idx] == blue {
			nums[idx], nums[bluePos] = nums[bluePos], nums[idx]
			bluePos -= 1
		} else {
			idx += 1
		}
	}
}
