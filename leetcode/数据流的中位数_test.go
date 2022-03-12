package leetcode

import (
	"container/heap"
	"fmt"
	"testing"
)

func Test_heapSort(t *testing.T) {
	h := &Heap{}
	for i := 0; i < 10; i++ {
		heap.Push(h, i)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(heap.Pop(h))
	}
}

func Test_findMedian(t *testing.T) {
	finder := &MedianFinder{big: new(Heap), small: new(Heap)}
	for i := 0; i < 10; i++ {
		finder.AddNum(i)
		fmt.Println(finder.FindMedian())
	}
}
