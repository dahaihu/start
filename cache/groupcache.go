package cache

import (
	"fmt"
	"github.com/golang/groupcache/lru"
)

func cacheExap() {
	cache := lru.New(2)
	cache.Add("bill", 20) // bill
	cache.Add("dable", 19) // dable -> bill
	v, ok := cache.Get("bill") // bill -> dable
	if ok {
		fmt.Printf("bill's age is %v\n", v)
	}

	cache.Add("cat", "18") // cat -> bill -> dable(evicted, length is 2)

	fmt.Printf("cache length is %d\n", cache.Len())
	_, ok = cache.Get("dable")
	if !ok {
		fmt.Printf("dable was evicted out\n")
	}
}
