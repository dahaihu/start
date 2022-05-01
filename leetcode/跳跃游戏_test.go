package leetcode

import (
	"fmt"
	"testing"
)

func Test_jump(t *testing.T) {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{2, 1}))
}

func Test_canJump(t *testing.T) {
	fmt.Println(canJump([]int{3, 2, 1, 0, 1}))
	fmt.Println(canJump([]int{3, 2, 1, 1, 1}))
}
