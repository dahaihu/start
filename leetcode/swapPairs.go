package leetcode

type ListNode struct {
	Val int
	Next *ListNode
}
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := &ListNode{}
	pre.Next = head
	res := pre
	for head != nil && head.Next != nil {
		next := head.Next
		nextHead := next.Next
		pre.Next = next
		next.Next = head
		head.Next = nextHead
		pre = head
		head = nextHead
	}
	return res.Next
}


func swapPairsRecur(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairsRecur(next.Next)
	next.Next = head
	return next
}
