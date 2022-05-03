package leetcode

import (
	"fmt"
	"testing"
)

func Test_trapRainWater(t *testing.T) {
	heightMap := [][]int{
		{1, 4, 3, 1, 3, 2},
		{3, 2, 1, 3, 2, 4},
		{2, 3, 3, 2, 3, 1},
	}
	fmt.Println(trapRainWater(heightMap))
}
