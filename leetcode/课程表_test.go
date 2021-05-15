package leetcode

import (
	"fmt"
	"testing"
)

func TestCanFinish(t *testing.T) {
	prerequisites := [][]int{{1, 2}, {2, 3}, {3, 4}}
	fmt.Println(canFinish(2, prerequisites))
}
