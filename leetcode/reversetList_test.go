package leetcode

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	a := ListNode{
		Val: 10,
	}
	a.Next = &ListNode{Val: 20}
	a.Next.Next = &ListNode{Val: 30}
	a.Next.Next.Next = &ListNode{Val: 40}
	a.Next.Next.Next.Next = &ListNode{Val: 50}
	res := reverseList(&a)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
