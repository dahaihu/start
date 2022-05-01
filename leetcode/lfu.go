package leetcode
//
//import (
//	"bytes"
//	"container/list"
//	"fmt"
//)
//
//type lfuNode struct {
//	key, value int
//	freqNode   *list.Element
//}
//
//func newLfuNode(key, value int, frequencyNode *list.Element) *lfuNode {
//	return &lfuNode{key: key, value: value, freqNode: frequencyNode}
//}
//
//type frequencyNode struct {
//	frequency int
//	items     map[int]*list.Element
//	itemsList *list.List
//}
//
//func newFrequencyNode(freq int) *frequencyNode {
//	return &frequencyNode{
//		frequency: freq,
//		itemsList: list.New(),
//		items:     make(map[int]*list.Element),
//	}
//}
//
//type LFUCache struct {
//	capacity int
//
//	list  *list.List
//	items map[int]*list.Element
//}
//
//func LFUConstructor(cap int) *LFUCache {
//	return &LFUCache{
//		capacity: cap,
//		list:     list.New(),
//		items:    make(map[int]*list.Element),
//	}
//}
//
//func (lfu *LFUCache) String() string {
//	var buffer bytes.Buffer
//	freqNode := lfu.list.Front()
//	var outerdep, innerdep int
//	for freqNode != nil {
//		outerdep++
//		innerdep = 0
//		realNode := freqNode.Value.(*frequencyNode)
//		buffer.WriteString(fmt.Sprintf("freq: [%d] ", realNode.frequency))
//		lfunode := realNode.itemsList.Front()
//		for lfunode != nil {
//			innerdep++
//			fmt.Printf("outdep %d, indep %d, lfunode type %T\n",
//				outerdep, innerdep, lfunode.Value)
//			reallfunode := lfunode.Value.(*lfuNode)
//			buffer.WriteString(fmt.Sprintf("->(%3d, %3d)",
//				reallfunode.key, reallfunode.value))
//			lfunode = lfunode.Next()
//		}
//		buffer.WriteString("\n")
//		freqNode = freqNode.Next()
//	}
//	return buffer.String()
//}
//
//func (lfu *LFUCache) evict() {
//	head := lfu.list.Front()
//
//	frequencyNode := head.Value.(*frequencyNode)
//	// 删除尾部节点
//	tail := frequencyNode.itemsList.Back()
//	node := tail.Value.(*lfuNode)
//	frequencyNode.itemsList.Remove(tail)
//	delete(frequencyNode.items, node.key)
//	delete(lfu.items, node.key)
//
//	if frequencyNode.itemsList.Len() == 0 {
//		lfu.list.Remove(head)
//	}
//}
//
//func (lfu *LFUCache) increment(ele *list.Element) {
//	realNode := ele.Value.(*lfuNode)
//	freqNode := realNode.freqNode
//	realFreqNode := freqNode.Value.(*frequencyNode)
//	// why ele.list is nil
//	realFreqNode.itemsList.Remove(ele)
//	delete(realFreqNode.items, realNode.key)
//
//	nextFreq := freqNode.Next()
//	if nextFreq == nil ||
//		nextFreq.Value.(*frequencyNode).frequency != realFreqNode.frequency+1 {
//		next := newFrequencyNode(realFreqNode.frequency + 1)
//		pushedEle := next.itemsList.PushFront(realNode)
//		next.items[realNode.key] = pushedEle
//		lfu.items[realNode.key] = pushedEle
//		realNode.freqNode = lfu.list.InsertAfter(next, freqNode)
//	} else {
//		realNode.freqNode = nextFreq
//		insertedFreq := nextFreq.Value.(*frequencyNode)
//		pushedEle := insertedFreq.itemsList.PushFront(realNode)
//		insertedFreq.items[realNode.key] = pushedEle
//		lfu.items[realNode.key] = pushedEle
//	}
//
//	if len(realFreqNode.items) == 0 {
//		lfu.list.Remove(freqNode)
//	}
//}
//
//func (lfu *LFUCache) Put(key, value int) {
//	if item, ok := lfu.items[key]; !ok {
//		if len(lfu.items) == lfu.capacity {
//			lfu.evict()
//		}
//		if len(lfu.items) == lfu.capacity {
//			return
//		}
//		head := lfu.list.Front()
//		if head == nil {
//			newFreqNode := newFrequencyNode(1)
//			insertedFreqNode := lfu.list.PushFront(newFreqNode)
//			lfunode := newLfuNode(key, value, insertedFreqNode)
//			lfuEle := newFreqNode.itemsList.PushFront(lfunode)
//			newFreqNode.items[key] = lfuEle
//			lfu.items[key] = lfuEle
//			return
//		}
//		frequencyHead := head.Value.(*frequencyNode)
//		if frequencyHead.frequency != 1 {
//			frequencyHead = newFrequencyNode(1)
//			head = lfu.list.PushFront(frequencyHead)
//		}
//		lfunode := newLfuNode(key, value, head)
//		lfuEle := frequencyHead.itemsList.PushFront(lfunode)
//		frequencyHead.items[key] = lfuEle
//		lfu.items[key] = lfuEle
//	} else {
//		lfu.increment(item)
//		realNode := item.Value.(*lfuNode)
//		realNode.value = value
//	}
//}
//
//func (lfu *LFUCache) Get(key int) int {
//	if lfu.capacity == 0 {
//		return -1
//	}
//	if item, ok := lfu.items[key]; ok {
//		lfu.increment(item)
//		node := item.Value.(*lfuNode)
//		return node.value
//	}
//	return -1
//}
