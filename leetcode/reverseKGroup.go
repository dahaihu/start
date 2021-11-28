package leetcode

// very interesting
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := new(ListNode)
	pre := dummy
	for {
		cur := head
		for i := 0; i < k; i++ {
			if cur == nil {
				pre.Next = head
				return dummy.Next
			}
			cur = cur.Next
		}

		segmentPre := head
		cur = head.Next
		head.Next = nil
		for i := 1; i < k; i++ {
			next := cur.Next
			cur.Next = segmentPre
			segmentPre, cur = cur, next
		}
		pre.Next = segmentPre
		pre = head
		head = cur
	}
}
