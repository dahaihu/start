package leetcode

import (
	"fmt"
	"testing"
)

func Test_next(t *testing.T) {
	root := &BinaryTreeNode{Val: 100}
	node1 := &BinaryTreeNode{Val: 1}
	node2 := &BinaryTreeNode{Val: 2}
	node3 := &BinaryTreeNode{Val: 3}
	node4 := &BinaryTreeNode{Val: 4}
	node5 := &BinaryTreeNode{Val: 5}
	node6 := &BinaryTreeNode{Val: 6}
	root.Left, node1.Parent = node1, root
	root.Right, node2.Parent = node2, root
	node1.Left, node3.Parent = node3, node1
	node1.Right, node4.Parent = node4, node1
	node2.Left, node5.Parent = node5, node2
	node2.Right, node6.Parent = node6, node2

	fmt.Println(node2.next())
}
