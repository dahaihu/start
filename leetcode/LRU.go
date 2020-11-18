package leetcode

import "fmt"

type Node struct {
	key, val   int
	prev, next *Node
}

type LRUCache struct {
	head, leaf       *Node
	capacity, length int
	nodes            map[int]*Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity, nodes: make(map[int]*Node)}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.nodes[key]; ok {
		if node == this.head {
			return node.val
		}
		prev := node.prev
		if node == this.leaf {
			node.prev = nil
			prev.next = nil
			this.leaf = prev
		} else {
			node.prev.next, node.next.prev = node.next, node.prev
		}
		node.next = this.head
		this.head.prev = node
		this.head = node
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// if exists, just change position
	if node, ok := this.nodes[key]; ok {
		node.val = value
		if node == this.head {
			return
		}
		prev := node.prev
		if node == this.leaf {
			node.prev = nil
			prev.next = nil
			this.leaf = prev
		} else {
			node.prev.next, node.next.prev = node.next, node.prev
		}
		node.next = this.head
		this.head.prev = node
		this.head = node
		return
	}
	// not exists, create a new node
	node := &Node{val: value, key: key}
	this.nodes[key] = node

	if this.length == 0 {
		this.head = node
		this.leaf = node
		this.length += 1
		return
	}

	// put node in the first position, and change the head
	node.next = this.head
	this.head.prev = node
	this.head = node

	// if length < capacity
	if this.length < this.capacity {
		this.length += 1
	} else {
		leaf := this.leaf
		prev := leaf.prev
		prev.next = nil
		leaf.prev = nil
		this.leaf = prev
		delete(this.nodes, leaf.key)
	}
}

func (this *LRUCache) print() {
	head := this.head
	for head != nil {
		if head.next == nil {
			fmt.Printf("(%d, %d)", head.key, head.val)
			break
		}
		fmt.Printf("(%d, %d) => ", head.key, head.val)
		head = head.next
	}
}
