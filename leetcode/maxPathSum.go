package leetcode

import "math"

type maxPathSumTreeNode struct {
	Val   int
	Left  *maxPathSumTreeNode
	Right *maxPathSumTreeNode
}

func maxPathMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxPathSum(root *maxPathSumTreeNode) int {
	res := math.MinInt64
	getMaxPath(root, &res)
	return res
}

func getMaxPath(node *maxPathSumTreeNode, res *int) int {
	if node == nil {
		return 0
	}
	left := maxPathMax(0, getMaxPath(node.Left, res))
	right := maxPathMax(0, getMaxPath(node.Right, res))
	if tmp := node.Val + left + right; tmp > *res {
		*res = tmp
	}
	return maxPathMax(left, right) + node.Val
}
