package leetcode

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0) == 4)
	fmt.Println(search([]int{3, 1}, 3))
}
