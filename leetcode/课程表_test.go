package leetcode

import (
	"fmt"
	"testing"
)

func TestCanFinish(t *testing.T) {
	prerequisites := [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}
	fmt.Println(canFinishBFS(5, prerequisites))
}
