package leetcode

import (
	"fmt"
	"testing"
)

func TestSortColors(t *testing.T) {
	original := []int{0, 1, 0}
	sortColors(original)
	fmt.Println(original)
}
