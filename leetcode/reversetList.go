package leetcode

func reverseListRec(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseListRec(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func reverseListRecursive(head *ListNode) (*ListNode, *ListNode) {
	if head.Next == nil {
		return head, head
	}
	next := head.Next
	// 第一个节点后面的节点
	head.Next = nil
	prev, res := reverseListRecursive(next)
	prev.Next = head

	return head, res
}
