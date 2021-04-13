package leetcode

import (
	"fmt"
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	a := []int{3, 2, 100, -1}
	fmt.Println(firstMissingPositive(a))
}
