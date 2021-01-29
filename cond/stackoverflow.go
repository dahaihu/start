package main

import (
	"fmt"
	"sync"
	"time"
)

var sharedRsc = make(map[string]interface{})

func stackoverflow() {
	var wg sync.WaitGroup
	wg.Add(2)
	m := sync.Mutex{}
	c := sync.NewCond(&m)

	go func() {
		// this go routine wait for changes to the sharedRsc
		c.L.Lock()
		defer c.L.Unlock()
		for len(sharedRsc) == 0 {
			c.Wait()
		}
		fmt.Println(sharedRsc["rsc1"])
		wg.Done()
	}()

	go func() {
		// this go routine wait for changes to the sharedRsc
		c.L.Lock()
		defer c.L.Unlock()
		for len(sharedRsc) == 0 {
			c.Wait()
		}
		fmt.Println(sharedRsc["rsc2"])
		wg.Done()
	}()
	time.Sleep(time.Second)
	// this one writes changes to sharedRsc
	c.L.Lock()
	sharedRsc["rsc1"] = "foo"
	sharedRsc["rsc2"] = "bar"
	c.L.Unlock()
	c.Broadcast()
	wg.Wait()
}
