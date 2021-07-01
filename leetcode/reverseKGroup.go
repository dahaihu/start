package leetcode

// very interesting
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	sentinel := dummy
	for {
		tmp := head
		for i := 0; i < k; i++ {
			if tmp == nil {
				sentinel.Next = head
				return dummy.Next
			}
			tmp = tmp.Next
		}
		pre, cur := head, head.Next
		pre.Next = nil
		for i := 1; i < k; i++ {
			tmp = cur.Next
			cur.Next = pre

			pre, cur = cur, tmp
		}
		sentinel.Next = pre
		sentinel = head
		head = cur
	}
}
