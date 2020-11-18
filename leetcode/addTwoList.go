package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return reverseListINAddTwoList(addList(l1, l2, 0))
}

func addList(l1, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}
	val := getNodeVal(l1) + getNodeVal(l2) + carry
	if val >= 10 {
		carry = 1
		val -= 10
	} else {
		carry = 0
	}
	node := &ListNode{}
	node.Val = val
	node.Next = addList(getNodeNext(l1), getNodeNext(l2), carry)
	return node
}
func getNodeNext(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}
	return node.Next
}
func getNodeVal(node *ListNode) int {
	if node == nil {
		return 0
	}
	return node.Val
}

func reverseListINAddTwoList(node *ListNode) *ListNode {
	var pre *ListNode
	for node != nil {
		nextNode := node.Next
		node.Next = pre
		pre = node
		node = nextNode
	}
	return pre
}
