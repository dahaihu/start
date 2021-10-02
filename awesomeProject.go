package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Data   int
	Left   *Node
	Right  *Node
	Bf     int
	Parent *Node
}

func buildTreeSelf(node *Node) (
	box []string, boxLength, rootStart, rootEnd int,
) {
	if node == nil {
		return nil, 0, 0, 0
	}
	leftBox, leftLen, leftRootStart, leftRootEnd := buildTree(node.Left)
	rightBox, rightLen, rightRootStart, rightRootEnd := buildTree(node.Right)

	rootStr := strconv.Itoa(node.Data)
	rootLen := len(rootStr)
	gapSize := rootLen

	var line1, line2 []string
	if leftLen > 0 {
		mid := (leftRootStart + leftRootEnd) / 2
		line1 = append(line1, strings.Repeat(" ", mid+2))
		line1 = append(line1, strings.Repeat("_", leftLen+1-(mid+2)))
		line2 = append(line2, strings.Repeat(" ", mid+1)+"/")
		line2 = append(line2, strings.Repeat(" ", leftLen+1-(mid+2)))
		rootStart = leftLen + 1
		gapSize++
	} else {
		rootStart = 0
	}
	rootEnd = rootStart + rootLen - 1
	line1 = append(line1, rootStr)
	line2 = append(line2, strings.Repeat(" ", rootLen))
	if rightLen > 0 {
		mid := (rightRootStart + rightRootEnd) / 2
		line1 = append(line1, strings.Repeat("_", mid))
		line1 = append(line1, strings.Repeat(" ", 1+rightLen-mid))
		line2 = append(line2, strings.Repeat(" ", mid)+"\\")
		line2 = append(line2, strings.Repeat(" ", 1+rightLen-mid-1))
		gapSize++
	}
	box = []string{strings.Join(line1, ""), strings.Join(line2, "")}
	boxLength = len(box[0])
	maxHeight := len(leftBox)
	if rightHeight := len(rightBox); rightHeight > maxHeight {
		maxHeight = rightHeight
	}
	gapStr := strings.Repeat(" ", gapSize)
	for i := 0; i < maxHeight; i++ {
		var leftLine, rightLine string
		if i < len(leftBox) {
			leftLine = leftBox[i]
		} else {
			leftLine = strings.Repeat(" ", leftLen)
		}
		if i < len(rightBox) {
			rightLine = rightBox[i]
		} else {
			rightLine = strings.Repeat(" ", rightLen)
		}
		box = append(box, leftLine+gapStr+rightLine)
	}
	return
}

func buildTree(node *Node) (
	box []string, boxLength, rootStart, rootEnd int,
) {
	if node == nil {
		return nil, 0, 0, 0
	}

	rootStr := strconv.Itoa(node.Data)
	rootWidth := len(rootStr)
	gapSize := len(rootStr)

	leftBox, leftBoxLength, leftRootStart, leftRootEnd := buildTree(node.Left)
	rightBox, rightBoxLength, rightRootStart, rightRootEnd := buildTree(node.Right)
	var (
		line1, line2 []string
	)
	if leftBoxLength > 0 {
		leftRootIndex := (leftRootStart+leftRootEnd)/2 + 1
		line1 = append(line1, strings.Repeat(" ", leftRootIndex+1))
		line1 = append(line1, strings.Repeat("_", leftBoxLength-leftRootIndex))
		line2 = append(line2, strings.Repeat(" ", leftRootIndex)+"/")
		line2 = append(line2, strings.Repeat(" ", leftBoxLength-leftRootIndex))
		rootStart = leftBoxLength + 1
		gapSize++
	} else {
		rootStart = 0
	}
	line1 = append(line1, rootStr)
	line2 = append(line2, strings.Repeat(" ", rootWidth))
	if rightBoxLength > 0 {
		rightRootIndex := (rightRootStart + rightRootEnd) / 2
		line1 = append(line1, strings.Repeat("_", rightRootIndex))
		line1 = append(line1, strings.Repeat(" ",
			rightBoxLength-rightRootIndex+1))
		line2 = append(line2, strings.Repeat(" ", rightRootIndex)+"\\")
		line2 = append(line2, strings.Repeat(" ",
			rightBoxLength-rightRootIndex))
		gapSize++
	}
	rootEnd = rootStart + rootWidth - 1
	gapStr := strings.Repeat(" ", gapSize)
	newBox := []string{strings.Join(line1, ""), strings.Join(line2, "")}
	childHeight := len(leftBox)
	if rightBoxHeight := len(rightBox); rightBoxHeight > childHeight {
		childHeight = rightBoxHeight
	}
	for i := 0; i < childHeight; i++ {
		var lline, rline string
		if i < len(leftBox) {
			lline = leftBox[i]
		} else {
			lline = strings.Repeat(" ", leftBoxLength)
		}
		if i < len(rightBox) {
			rline = rightBox[i]
		} else {
			rline = strings.Repeat(" ", rightBoxLength)
		}
		newBox = append(newBox, lline+gapStr+rline)
	}
	return newBox, len(newBox[0]), rootStart, rootEnd
}

//平衡二叉树，也叫 AVL 树（平衡二叉树作者的名字首字母），是自平衡的二叉查找树，
//要求每个节点的左子树和右子树的高度差至多等于 1，
//这个高度（深度）差的值叫做平衡因子 BF，也就是说 BF 的值不能大于1，
//距离插入节点最近的，且平衡因子绝对值大于 1 的节点为根的子树，叫做最小不平衡子树，
//一旦出现最小不平衡子树时，就进行左旋、右旋或双旋处理，以保持自身始终平衡
//算法复杂度：O(logn)
type AvlTree struct {
	Tree *Node
}

const (
	LH = 1
	EH = 0
	RH = -1
)

//中序遍历
func (t *AvlTree) MidOrderTraverse(tree *Node) {
	if tree == nil {
		return
	}

	t.MidOrderTraverse(tree.Left)
	// fmt.Println(tree.Data)
	fmt.Println(tree)
	t.MidOrderTraverse(tree.Right)
}

func (t *AvlTree) Insert(data int) {
	t.InsertNode(data)
	t.Reset() //将指针恢复到整棵树的根节点处
}

func (t *AvlTree) Reset() {
	tree := t.Tree
	for tree.Parent != nil {
		tree = tree.Parent
	}
	t.Tree = tree
}

func (t *AvlTree) InsertNode(data int) bool {
	if t.Tree == nil {
		t.Tree = &Node{Data: data, Bf: EH}
		return true
	}
	tree := t.Tree
	if data < tree.Data {
		t.Tree = tree.Left
		if !t.InsertNode(data) {
			return false
		} else {
			if t.Tree.Parent == nil {
				t.Tree.Parent = tree
			}
			if tree.Left == nil {
				tree.Left = t.Tree
			}

			switch tree.Bf {
			case LH:
				t.LeftBalance(tree)
				return false
			case EH:
				tree.Bf = LH
				t.Tree = tree
				return true
			case RH:
				tree.Bf = EH
				return false
			}
		}
	} else if data > tree.Data {
		t.Tree = tree.Right
		if !t.InsertNode(data) {
			return false
		} else {
			if t.Tree.Parent == nil {
				t.Tree.Parent = tree
			}
			if tree.Right == nil {
				tree.Right = t.Tree
			}

			switch tree.Bf {
			case RH:
				t.RightBalance(tree)
				return false
			case EH:
				tree.Bf = RH
				t.Tree = tree
				return true
			case LH:
				tree.Bf = EH
				return false
			}
		}
	}
	return true
}

func (t *AvlTree) LeftBalance(tree *Node) {
	subTree := tree.Left
	if subTree != nil {
		switch subTree.Bf {
		case LH:
			// 新插入节点在左子节点的左子树上要做右单旋处理
			tree.Bf = EH
			subTree.Bf = EH
			t.RightRotate(tree)
		case RH:
			// 新插入节点在左子节点的右子树上要做双旋处理
			subTree_r := subTree.Right
			if subTree_r != nil {
				switch subTree_r.Bf {
				case LH:
					tree.Bf = RH
					subTree.Bf = EH
				case RH:
					tree.Bf = EH
					subTree.Bf = LH
				}
				subTree_r.Bf = EH
				t.LeftRotate(subTree)
				t.RightRotate(tree)
			}

		}
	}
}

func (t *AvlTree) RightBalance(tree *Node) {
	subTree := tree.Right
	if subTree != nil {
		switch subTree.Bf {
		case RH:
			//新插入节点在右子节点的右子树上要做左单旋处理
			tree.Bf = EH
			subTree.Bf = EH
			t.LeftRotate(tree)
		case LH:
			//新插入节点在右子节点的左子树上要做双旋处理
			subTree_l := subTree.Left
			if subTree_l != nil {
				switch subTree_l.Bf {
				case LH:
					tree.Bf = EH
					subTree.Bf = RH
				case RH:
					tree.Bf = LH
					subTree.Bf = EH
				}
				subTree_l.Bf = EH
				t.RightRotate(subTree)
				t.LeftRotate(tree)
			}

		}
	}
}

//右单旋
func (t *AvlTree) RightRotate(tree *Node) {
	subTree := tree.Left
	isLeft := false
	if tree.Parent != nil {
		subTree.Parent = tree.Parent //更新新子树的父节点
		if tree.Parent.Left == tree {
			isLeft = true
		}
	} else {
		subTree.Parent = nil
	}
	tree.Left = subTree.Right //原来左节点的右子树挂到老的根节点的左子树
	if subTree.Right != nil {
		subTree.Right.Parent = tree
	}
	tree.Parent = subTree //原来的左节点变成老的根节点的父节点
	subTree.Right = tree  //原来的根节点变成原来左节点的右子树
	tree = subTree
	if tree.Parent == nil { //旋转的是整棵树的根节点
		t.Tree = tree
	} else {
		if isLeft { //更新老的子树根节点父节点指针指向新的根节点
			tree.Parent.Left = tree
		} else {
			tree.Parent.Right = tree
		}
	}
}

//左单旋
func (t *AvlTree) LeftRotate(tree *Node) {
	subTree := tree.Right
	isLeft := false
	if tree.Parent != nil {
		subTree.Parent = tree.Parent
		if tree.Parent.Left == tree {
			isLeft = true
		}
	} else {
		subTree.Parent = nil
	}
	tree.Right = subTree.Left
	if subTree.Left != nil {
		subTree.Left.Parent = tree
	}
	tree.Parent = subTree
	subTree.Left = tree
	tree = subTree
	if tree.Parent == nil {
		t.Tree = tree
	} else {
		if isLeft {
			tree.Parent.Left = tree
		} else {
			tree.Parent.Right = tree
		}
	}
}

func main() {
	avlTree := &AvlTree{}

	//右单旋测试
	avlTree.Insert(1)
	avlTree.Insert(2)
	avlTree.Insert(3)
	avlTree.Insert(4)
	avlTree.Insert(5)
	avlTree.Insert(6)
	//avlTree.Insert(7)
	box, _, _, _ := buildTreeSelf(avlTree.Tree)
	for _, line := range box {
		fmt.Println(line)
	}
	//avlTree.Insert(66)
	//avlTree.Insert(83)
	//avlTree.Insert(110)
	//avlTree.Insert(130)
	//fmt.Println(avlTree)
}
