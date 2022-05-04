package leetcode

//
//import (
//	"bytes"
//	"fmt"
//)
//
//type lruNode struct {
//	prev, next *lruNode
//	key, value int
//}
//
//type list struct {
//	headDummy, tailDummy *lruNode
//}
//
//func (l *list) saveToHead(node *lruNode) {
//	next := l.headDummy.next
//	// move node to head
//	l.headDummy.next = node
//	node.prev = l.headDummy
//
//	node.next = next
//	next.prev = node
//}
//
//func (l *list) remove(node *lruNode) {
//	node.prev.next, node.next.prev = node.next, node.prev
//	node.next = nil
//	node.next = nil
//}
//
//func (l *list) tail() *lruNode {
//	if l.headDummy.next != l.tailDummy {
//		return l.tailDummy.prev
//	}
//	return nil
//}
//
//func (l *list) head() *lruNode {
//	if l.headDummy.next != l.tailDummy {
//		return l.headDummy.next
//	}
//	return nil
//}
//
//func (l *list) String() string {
//	buf := new(bytes.Buffer)
//	for node := l.head(); node != l.tailDummy; node = node.next {
//		buf.WriteString(fmt.Sprintf("-> %d(%d)", node.key, node.value))
//	}
//	return buf.String()
//}
//
//type LRUCache struct {
//	list             *list
//	nodes            map[int]*lruNode
//	capacity, length int
//}
//
//func Constructor(capacity int) LRUCache {
//	head, tail := new(lruNode), new(lruNode)
//	head.next, tail.next = tail, head
//	return LRUCache{
//		list:     &list{headDummy: head, tailDummy: tail},
//		nodes:    make(map[int]*lruNode),
//		capacity: capacity,
//	}
//}
//
//func (this *LRUCache) moveToHead(node *lruNode) {
//	this.list.remove(node)
//	this.list.saveToHead(node)
//}
//
//func (this *LRUCache) Get(key int) int {
//	if node, ok := this.nodes[key]; ok {
//		this.moveToHead(node)
//		return node.value
//	}
//	return -1
//}
//
//func (this *LRUCache) add(node *lruNode) {
//	this.nodes[node.key] = node
//	this.list.saveToHead(node)
//}
//
//func (this *LRUCache) remove(node *lruNode) {
//	delete(this.nodes, node.key)
//	this.list.remove(node)
//}
//
//func (this *LRUCache) Put(key int, value int) {
//	if node, ok := this.nodes[key]; ok {
//		node.value = value
//		this.moveToHead(node)
//		return
//	}
//	node := &lruNode{key: key, value: value}
//	this.add(node)
//	this.length += 1
//	if this.length > this.capacity {
//		node := this.list.tail()
//		this.remove(node)
//		this.length -= 1
//		return
//	}
//}
