package leetcode

import (
	"fmt"
	"testing"
)

func TestLRUCache_Get(t *testing.T) {
	lru := Constructor(1)
	lru.Put(2, 1)
	fmt.Println(lru.length, len(lru.nodes))

	lru.Get(2)
	fmt.Println(lru.length, len(lru.nodes))

	lru.Put(3, 2)
	lru.Get(2)
	lru.Get(3)
	fmt.Println(lru.length, len(lru.nodes))
	lru.print()

}
