package leetcode

import (
	"fmt"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	root := TreeNode{Val: 1}
	rightChild := TreeNode{Val: 2}
	leftGrandchildren := TreeNode{Val: 3}
	root.Right = &rightChild
	rightChild.Left = &leftGrandchildren
	fmt.Println(inorderTraversal(&root))

}
