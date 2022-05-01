package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}) == 4)
	fmt.Println(standardLongestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}) == 9)
}
