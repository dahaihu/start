package main

import (
	"fmt"
	"strings"
)

func buildTree(node *Node) (
	box []string, boxLength, rootStart, rootEnd int,
) {
	if node == nil {
		return nil, 0, 0, 0
	}

	rootStr := fmt.Sprintf("%d(%d)", node.Data, node.Height)
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

type Node struct {
	Left   *Node
	Right  *Node
	Data   int
	Height int
}

func (n *Node) height() int {
	if n == nil {
		return 0
	}
	return n.Height
}

func (n *Node) updateHeight() {
	n.Height = max(n.Left.height(), n.Right.height()) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (n *Node) leftRotate() *Node {
	head := n.Right
	n.Right = head.Left
	head.Left = n
	n.updateHeight()
	head.updateHeight()
	return head
}

func (n *Node) rightRotate() *Node {
	head := n.Left
	n.Left = head.Right
	head.Right = n
	n.updateHeight()
	head.updateHeight()
	return head
}

func (n *Node) adjust() *Node {
	nodeStatus := n.balanceStatus()
	switch {
	case nodeStatus == 2:
		if n.Left.balanceStatus() == 1 {
			return n.rightRotate()
		} else {
			n.Left = n.Left.leftRotate()
			return n.rightRotate()
		}
	case nodeStatus == -2:
		if n.Right.balanceStatus() == -1 {
			return n.leftRotate()
		} else {
			n.Right = n.Right.rightRotate()
			return n.leftRotate()
		}
	}
	return n
}

func (n *Node) balanceStatus() int {
	return n.Left.height() - n.Right.height()
}

func (n *Node) Insert(data int) (node *Node, ok bool) {
	if n == nil {
		return NewNode(data), true
	}
	if data == n.Data {
		return nil, false
	} else if data < n.Data {
		n.Left, ok = n.Left.Insert(data)
	} else {
		n.Right, ok = n.Right.Insert(data)
	}
	n.updateHeight()
	return n.adjust(), ok
}

func (n *Node) leftMax() *Node {
	cur := n.Left
	if cur == nil {
		return cur
	}
	for cur.Right != nil {
		cur = cur.Right
	}
	return cur
}

func (n *Node) Delete(data int) (node *Node, ok bool) {
	if n == nil {
		return nil, false
	}
	if data == n.Data {
		deletedNode := n.leftMax()
		if deletedNode == nil {
			return n.Right, true
		}
		n.Data = deletedNode.Data
		n.Left, ok = n.Left.Delete(n.Data)
	} else if data < n.Data {
		n.Left, ok = n.Left.Delete(data)
	} else {
		n.Right, ok = n.Right.Delete(data)
	}
	n.updateHeight()
	return n.adjust(), ok
}

func NewNode(data int) *Node {
	return &Node{Data: data, Height: 1}
}

func (n *Node) print() {
	box, _, _, _ := buildTree(n)
	for _, line := range box {
		fmt.Println(line)
	}
}

func main() {
	// 速度的搞定问题，解决问题
	n := NewNode(-1)
	n, _ = n.Insert(1)
	n.print()
	n, _ = n.Insert(11)
	n.print()
	n, _ = n.Insert(12)
	n.print()
	n, _ = n.Insert(9)
	n.print()
	n, _ = n.Delete(11)
	n.print()
	//n, _ = n.Insert(15)
	//n.print()
	//n, _ = n.Insert(17)
	//n.print()
	//n, _ = n.Insert(14)
	//n.print()
	//n, _ = n.Insert(16)
	//n.print()
	//n, _ = n.Insert(18)
	//n.print()
	//n, _ = n.Insert(19)
	//n.print()
	//n, _ = n.Insert(20)
	//n.print()
	//n, _ = n.Insert(21)
	//n.print()
	//n, _ = n.Delete(18)
	//n.print()
	//n, _ = n.Delete(19)
	//n.print()
	//n, _ = n.Delete(17)
	//n.print()
	//n, _ = n.Delete(16)
	//n.print()
	//fmt.Println(n.Data, n)
	//fmt.Println(n.Left)
	//fmt.Println(n.Right)
}
