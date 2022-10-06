package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func TestMerge(t *testing.T) {
	fmt.Println(merge(NewIntervals([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})))
	fmt.Println(merge(NewIntervals([][]int{{1, 4}, {2, 3}})))
}

func TestSort(t *testing.T) {
	values := []int{3, 1, 4, 1, 10}
	sort.Slice(values, func(i, j int) bool { return values[i] <= values[j] })
	fmt.Println(values)
}
