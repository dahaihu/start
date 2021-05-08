package leetcode

import "fmt"

type lruNode struct {
	prev, next *lruNode
	key, val   int
}

type LRUCache struct {
	m                    map[int]*lruNode
	length, cap          int
	dummyHead, dummyTail *lruNode
}

func Constructor(capacity int) LRUCache {
	head, tail := &lruNode{}, &lruNode{}
	head.next, tail.prev = tail, head
	return LRUCache{
		m:         make(map[int]*lruNode),
		length:    0,
		cap:       capacity,
		dummyHead: head,
		dummyTail: tail,
	}
}
func (this *LRUCache) moveToHead(node *lruNode) {
	if node == this.dummyHead.next {
		return
	}
	// ! node < !
	node.prev.next, node.next.prev = node.next, node.prev

	head := this.dummyHead.next
	// dummyHead -> node-> head
	this.dummyHead.next, node.next = node, head
	// dummyHead <- node <- head
	head.prev, node.prev = node, this.dummyHead
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		this.moveToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		this.moveToHead(node)
		node.val = value
		return
	}

	head := this.dummyHead.next
	node := &lruNode{key: key, val: value}
	// add data to map
	this.m[key] = node

	// add data to linked list
	// dummyHead -> node-> head
	this.dummyHead.next, node.next = node, head
	// dummyHead <- node <- head
	head.prev, node.prev = node, this.dummyHead

	if this.length < this.cap {
		this.length += 1
		return
	}

	// remove tail
	tail := this.dummyTail.prev
	// delete data from map
	delete(this.m, tail.key)
	// delete data from linked map
	tail.prev.next, tail.next.prev = tail.next, tail.prev

	tail.prev = nil
	tail.next = nil
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func (this *LRUCache) print() {
	head := this.dummyHead.next
	for {
		if head.next == this.dummyTail {
			fmt.Printf("(%d, %d)", head.key, head.val)
			break
		}
		fmt.Printf("(%d, %d) => ", head.key, head.val)
		head = head.next
	}
}
