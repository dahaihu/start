package leetcode

import (
	"fmt"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	one, two, three, four, five, six, seven := &ListNode{Val:1},
	&ListNode{Val:2},
	&ListNode{Val:3},
	&ListNode{Val:4},
	&ListNode{Val:5},
	&ListNode{Val:6},
	&ListNode{Val:7}
	one.Next = two
	two.Next = three
	three.Next = four
	four.Next = five
	five.Next = six
	six.Next = seven


	res := reverseKGroup(one, 2)
	for res != nil {
		fmt.Printf("%d -> ", res.Val)
		res = res.Next
	}
}
