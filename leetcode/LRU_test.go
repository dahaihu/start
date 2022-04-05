package leetcode

import (
	"fmt"
	"testing"
)

func TestLRUCache_Get(t *testing.T) {
	lru := Constructor(2)
	lru.Put(1, 1)
	fmt.Println(lru.list)

	lru.Put(2, 2)
	fmt.Println(lru.list)

	fmt.Println(lru.Get(1))
	fmt.Println(lru.list)

	lru.Put(3, 3)
	fmt.Println(lru.list)

	fmt.Println(lru.Get(2))
	lru.Put(4, 4)
	fmt.Println(lru.list)

	fmt.Println(lru.Get(1))
	fmt.Println(lru.list)

	fmt.Println(lru.Get(3))
	fmt.Println(lru.list)

	fmt.Println(lru.Get(4))
	fmt.Println(lru.list)

	fmt.Println(lru.list)
}
