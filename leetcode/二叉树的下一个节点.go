package leetcode

type BinaryTreeNode struct {
	Left, Right, Parent *BinaryTreeNode
	Val                 int
}

func (n *BinaryTreeNode) next() *BinaryTreeNode {
	if n.Right != nil {
		cur := n.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		return cur
	} else if n.Parent.Left == n {
		return n.Parent
	} else if n.Parent != nil && n.Parent.Right == n {
		cur := n.Parent
		for cur.Parent != nil && cur.Parent.Right == cur {
			cur = cur.Parent
		}
		if cur.Parent == nil {
			return nil
		} else {
			return cur.Parent
		}
	}
	return nil
}
