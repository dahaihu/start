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
		cur = head
		var tmpPre *ListNode
		for i := 0; i < k; i++ {
			next := cur.Next
			cur.Next = tmpPre
			tmpPre = cur
			cur = next
		}
		pre.Next = tmpPre
		pre = head
		head = cur
	}
}
