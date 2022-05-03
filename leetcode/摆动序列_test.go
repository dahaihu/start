package leetcode

import (
	"fmt"
	"testing"
)

func Test_wiggleMaxLength(t *testing.T) {
	fmt.Println(wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}))
	fmt.Println(wiggleMaxLength([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println(wiggleMaxLength([]int{2, 2, 3, 3}))
	fmt.Println()
}

func TestWinggleMaxLength(t *testing.T) {
	fmt.Println(wiggleMaxLength([]int{1, 5, 3, 2, 2, 10}))
}
