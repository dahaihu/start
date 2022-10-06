package leetcode

import (
	"fmt"
	"testing"
)

func TestSortColors(t *testing.T) {
	original := []int{2, 0, 2, 1, 1, 0}
	sortColors(original)
	fmt.Println(original)
}
