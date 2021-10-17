package leetcode

import (
	"fmt"
	"testing"
)

func TestLFU(t *testing.T) {
	lfu := Constructor(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	fmt.Println(lfu.Get(1))
	lfu.Put(3, 3)
	//fmt.Println(lfu.Get(1))
	//lfu.Put(3, 3)
	fmt.Println(lfu.Get(2))
	fmt.Println(lfu.Get(3))
	fmt.Println(lfu)
	fmt.Println(lfu.items)
}
