package leetcode

import (
	"fmt"
	"testing"
)

func TestFindKth(t *testing.T) {
	// how to get out of the closed quan
	fmt.Println(FindMedianSortedArraysUsingSplit([]int{1, 2}, []int{3, 4}))
}

func TestDividedSortedArrays(t *testing.T) {
	fmt.Println(splitTwoSortedArrays([]int{1}, []int{2,3,4,5}))
}

type TestNode struct {
	key, val int
}

func TestNodeKey(t *testing.T) {
	node1 := TestNode{1, 2}
	dict := make(map[TestNode]bool)
	dict[node1] = true
	node2 := TestNode{1, 2}
	val, ok := dict[node2]
	fmt.Printf("val is %v, ok is %v\n", val, ok)
}
