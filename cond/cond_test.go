package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"testing"
	"time"
)

func listen(name string, a map[string]int, c *sync.Cond) {
	time.Sleep(time.Microsecond * 100)
	c.L.Lock()
	c.Wait()
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

	// 如果先 wait 呢？？？
	// listener 1
	go listen("lis1", age, cond)

	// listener 2
	go listen("lis2", age, cond)

	// broadcast
	// 如果先 broadcast 呢？？？？
	go broadcast("b1", age, cond)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
