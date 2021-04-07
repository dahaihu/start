package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	root := maxPathSumTreeNode{Val: 100}
	left := maxPathSumTreeNode{Val: 100}
	right := maxPathSumTreeNode{Val: 100}
	root.Left = &left
	root.Right = &right

	fmt.Println(maxPathSum(&root))
}


func TestString(t *testing.T) {
	items := strings.Split("1,2,3", "")
	fmt.Println(len(items) == 5)
}