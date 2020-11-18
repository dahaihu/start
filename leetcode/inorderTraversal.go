package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var _traversal func(*TreeNode) []int
	_traversal = func(node *TreeNode) []int {
		if node == nil {
			return make([]int, 0)
		}
		leftRes := _traversal(node.Left)
		rightRes := _traversal(node.Right)
		res := make([]int, len(leftRes)+len(rightRes)+1)
		copy(res[:len(leftRes)], leftRes)
		copy(res[len(res)-len(rightRes):], rightRes)
		res[len(leftRes)] = node.Val
		return res

	}
	return _traversal(root)
}

func inorderTraversalStandard(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}
	leftRes := inorderTraversal(root.Left)
	rightRes := inorderTraversal(root.Right)
	res := make([]int, len(leftRes)+len(rightRes)+1)
	copy(res[:len(leftRes)], leftRes)
	copy(res[len(res)-len(rightRes):], rightRes)
	res[len(leftRes)] = root.Val
	return res
}

func inorderTraversalBest(root *TreeNode) []int {
	res, mark := make([]int, 0), make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(mark) != 0 {
		if cur == nil {
			cur = mark[len(mark)-1]
			mark = mark[:len(mark)-1]

			res = append(res, cur.Val)
			cur = cur.Right

		} else {
			mark = append(mark, cur)
			cur = cur.Left
		}
	}
	return res
}
