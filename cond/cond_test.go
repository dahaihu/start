package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"testing"
)

func listen(name string, a map[string]int, c *sync.Cond) {
	c.L.Lock()
	for len(a) == 0 {
		c.Wait()
	}
	fmt.Printf("%s age: %d\n", name, a["T"])
	c.L.Unlock()
}

func broadcast(name string, a map[string]int, c *sync.Cond) {
	c.L.Lock()
	a["T"] = 25
	c.L.Unlock()
	c.Broadcast()
}

func TestCond(t *testing.T) {
	var age = make(map[string]int)

	m := sync.Mutex{}
	cond := sync.NewCond(&m)

	// listener 1
	go listen("lis1", age, cond)

	// listener 2
	go listen("lis2", age, cond)

	// broadcast
	go broadcast("b1", age, cond)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
