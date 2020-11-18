package leetcode

import (
	"fmt"
	"testing"
)

func TestSumNumbers(t *testing.T) {
	root := TreeNode{Val: 4}
	left := TreeNode{Val: 9}
	left.Left = &TreeNode{Val: 5}
	left.Right = &TreeNode{Val: 1}
	root.Left = &left
	root.Right = &TreeNode{Val: 0}
	fmt.Println(sumNumbers(&root) == 1026)
}