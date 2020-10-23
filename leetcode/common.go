package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func initiateList(values []int) *ListNode {
	firstNode := &ListNode{Val: values[0]}
	cur := firstNode
	for i := 1; i < len(values); i++ {
		node := &ListNode{Val: values[i]}
		cur.Next = node
		cur = node
	}
	return firstNode
}

func traverseList(node *ListNode) {
	for node != nil {
		if node.Next == nil {
			fmt.Printf("%d", node.Val)
			return
		}
		fmt.Printf("%d => ", node.Val)
		node = node.Next
	}
	fmt.Printf("\n")
}
