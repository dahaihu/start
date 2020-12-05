package leetcode

import (
	"fmt"
	"testing"
)

func TestLRUCache_Get(t *testing.T) {
	lru := Constructor(2)
	lru.Put(1, 1)
	fmt.Println(lru.length, len(lru.m))

	lru.Put(2, 2)
	lru.Get(1)
	//lru.Get(3)
	lru.Put(3, 3)
	lru.Get(2)
	lru.Put(4, 4)
	lru.Get(1)
	lru.Get(3)
	lru.Get(4)

	fmt.Println(lru.length, len(lru.m))
	lru.print()

}
