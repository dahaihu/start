package leetcode

// very interesting
func reverseKGroup(head *ListNode, k int) *ListNode {
	var left, node, dummy, sentinel *ListNode
	sentinel = &ListNode{}
	dummy = sentinel
	for head != nil {
		node = head
		for i := 0; i < k; i++ {
			if node == nil {
				sentinel.Next = head
				return dummy.Next
			}
			node = node.Next
		}
		left, node = head, nil
		for i := 0; i < k; i++ {
			next := head.Next
			head.Next = node
			node = head
			head = next
		}
		sentinel.Next = node
		sentinel = left
	}
	return dummy.Next
}