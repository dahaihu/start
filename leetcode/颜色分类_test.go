package leetcode

import (
	"fmt"
	"testing"
)

func TestSortColors(t *testing.T) {
	original := []int{0, 1, 2, 1}
	sortColors(original)
	fmt.Println(original)
}
