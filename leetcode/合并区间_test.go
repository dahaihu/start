package leetcode

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {2, 3}}))
}
