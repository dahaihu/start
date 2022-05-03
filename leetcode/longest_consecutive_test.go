package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	mark := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(mark))
}
