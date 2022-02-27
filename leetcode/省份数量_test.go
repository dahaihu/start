package leetcode

import (
	"fmt"
	"testing"
)

func TestFindCircleNum(t *testing.T) {
	fmt.Println(findCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}))
}
