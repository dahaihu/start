package leetcode

const (
	red   = 0
	white = 1
	blue  = 2
)

func sortColors(nums []int) {
	nextRed, nextBlue := 0, len(nums)-1
	for i := 0; i <= nextBlue; {
		switch nums[i] {
		case red:
			nums[i], nums[nextRed] = nums[nextRed], nums[i]
			nextRed += 1
			i += 1
		case white:
			i += 1
		case blue:
			nums[i], nums[nextBlue] = nums[nextBlue], nums[i]
			nextBlue -= 1
		}
	}
}
