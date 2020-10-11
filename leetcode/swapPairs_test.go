package leetcode

import (
	"fmt"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	head := &ListNode{Val: 0}
	pre := head
	for i := 1; i < 5; i++ {
		node := ListNode{Val: i}
		pre.Next = &node
		pre = &node
	}
	//for head != nil {
	//	fmt.Println("val is ", head.Val)
	//	head = head.Next
	//}
	// think myself as a sb
	res := swapPairsRecur(head)
	for res != nil {
		fmt.Println("val is ", res.Val)
		res = res.Next
	}
	fmt.Println('9' - '1')

}
