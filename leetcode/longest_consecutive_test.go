package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	mark := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1, 100}
	fmt.Println(longestConsecutiveUsingSet(mark))
	sort.Slice(mark, func(i, j int) bool { return mark[i] <= mark[j] })
	fmt.Println(mark)
}
