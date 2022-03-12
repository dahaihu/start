package leetcode

import (
	"container/heap"
	"sort"
)

type Heap struct {
	sort.IntSlice
}

func (h *Heap) Push(ele interface{}) {
	h.IntSlice = append(h.IntSlice, ele.(int))
}

func (h *Heap) Pop() interface{} {
	length := len(h.IntSlice)
	ele := h.IntSlice[length-1]
	h.IntSlice = h.IntSlice[:length-1]
	return ele
}

type MedianFinder struct {
	big, small *Heap
}

func (this *MedianFinder) AddNum(num int) {
	if this.big.Len() == 0 || num > this.big.IntSlice[0] {
		heap.Push(this.big, num)
		if this.big.Len() > this.small.Len()+1 {
			ele := heap.Pop(this.big).(int)
			heap.Push(this.small, -ele)
		}
	} else {
		heap.Push(this.small, -num)
		if this.small.Len() > this.big.Len() {
			ele := heap.Pop(this.small).(int)
			heap.Push(this.big, -ele)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	length := this.small.Len() + this.big.Len()
	if length%2 == 0 {
		return float64(this.big.IntSlice[0]-this.small.IntSlice[0]) / 2
	}
	return float64(this.big.IntSlice[0])
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
