package leetcode

import "fmt"

type Node struct {
	key, val   int
	prev, next *Node
}
type LRUCache struct {
	// root.prev is head
	// root.next is tail
	head, tail        *Node
	m           map[int]*Node
	length, cap int
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{length: 0, cap: capacity, head: head, tail: tail, m: make(map[int]*Node)}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		node.next.prev, node.prev.next = node.prev, node.next

		tmp := this.head.next

		tmp.prev = node
		node.next = tmp

		this.head.next = node
		node.prev = this.head

		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		node.val = value
		node.next.prev, node.prev.next = node.prev, node.next

		tmp := this.head.next

		tmp.prev = node
		node.next = tmp

		this.head.next = node
		node.prev = this.head

		return
	}
	node := &Node{key: key, val: value}
	realHead := this.head.next
	realHead.prev = node
	node.next = realHead

	this.head.next = node
	node.prev = this.head

	this.m[key] = node

	if this.length < this.cap {
		this.length += 1
		return
	}

	realTail := this.tail.prev

	realTail.next.prev, realTail.prev.next = realTail.prev, realTail.next

	realTail.next = nil
	realTail.prev = nil
	delete(this.m, realTail.key)
	return

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func (this *LRUCache) print() {
	head := this.head.next
	for {
		if head.next == this.tail {
			fmt.Printf("(%d, %d)", head.key, head.val)
			break
		}
		fmt.Printf("(%d, %d) => ", head.key, head.val)
		head = head.next
	}
}
