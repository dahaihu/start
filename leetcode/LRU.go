package leetcode

import "fmt"

type Node struct {
	key, val   int
	prev, next *Node
}

type LRUCache struct {
	head, tail        *Node
	m           map[int]*Node
	length, cap int
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{head: head, tail: tail, length: 0, cap: capacity, m: make(map[int]*Node)}
}

func (lru *LRUCache) adjustPlace(node *Node) {
	node.prev.next, node.next.prev = node.next, node.prev

	head := lru.head.next
	head.prev, node.next = node, head
	node.prev, lru.head.next = lru.head, node
}

func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.m[key]; ok {
		// head
		if node == lru.head.next {
			return node.val
		}
		lru.adjustPlace(node)
		return node.val
	}
	return -1
}


func (lru *LRUCache) Put(key int, val int) {
	if node, ok := lru.m[key]; ok {
		node.val = val
		lru.adjustPlace(node)
		return
	}

	node := &Node{key: key, val: val}
	// add to link head
	head := lru.head.next
	head.prev, node.next = node, head
	node.prev, lru.head.next = lru.head, node

	// add to m
	lru.m[key] = node


	if lru.length < lru.cap {
		lru.length += 1
		return
	}

	// delete tail
	tail := lru.tail.prev
	tail.next.prev, tail.prev.next = tail.prev, tail.next

	tail.prev = nil
	tail.next = nil

	delete(lru.m, tail.key)

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
