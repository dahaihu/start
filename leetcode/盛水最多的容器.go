package leetcode

func MaxArea(height []int) int {
	var maxWater int
	left, right := 0, len(height)-1
	for left < right {
		waterHeight := heightMin(height[left], height[right])
		maxWater = heightMax(maxWater, waterHeight*(right-left))
		if height[left] <= height[right] {
			left += 1
		} else {
			right -= 1
		}
	}
	return maxWater
}

func heightMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func heightMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
