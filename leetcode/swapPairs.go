package leetcode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := &ListNode{}
	dummy := pre
	for head != nil && head.Next != nil {
		next, nextHead := head.Next, head.Next.Next
		pre.Next = next
		next.Next = head
		head.Next = nextHead
		pre = head
		head = nextHead
	}
	return dummy.Next
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
