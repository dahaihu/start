package leetcode

import (
	"fmt"
	"testing"
)

func TestCheck(t *testing.T) {
	fmt.Println(check([][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}, 13, 10, 3))
	fmt.Println(kthSmallest([][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}, 1))
}
