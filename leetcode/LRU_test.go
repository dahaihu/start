package leetcode

import (
	"fmt"
	"testing"
)

func TestLRUCache_Get(t *testing.T) {
	lru := Constructor(2)
	lru.Put(1, 1)
	fmt.Println(lru.m, len(lru.m))

	lru.Put(2, 2)

	fmt.Println(lru.m)

}
