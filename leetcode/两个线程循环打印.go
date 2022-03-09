package leetcode

import (
	"fmt"
	"sync"
)

type User struct {
	Name string `json:"name,omitempty"`
	Age  int64  `json:"age,omitempty"`
}

func twoThreadPrint(times int) {
	signal := struct{}{}
	one, two := make(chan struct{}), make(chan struct{})
	go func() {
		for {
			_, ok := <-one
			if !ok {
				break
			}
			fmt.Println(1)
			two <- signal
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < times; i++ {
			<-two
			fmt.Println(2)
			if i == times-1 {
				close(one)
				break
			}
			one <- signal
		}
	}()
	one <- signal
	wg.Wait()
}

func twoThreadPrintUsingSignal() {
	var num int
	cond := sync.NewCond(new(sync.Mutex))
	count := 10
	for i := 0; i < count; i++ {
		go func(idx int) {
			for {
				cond.L.Lock()
				for idx != num {
					cond.Wait()
				}
				fmt.Println(idx)
				num = (idx + 1) % count
				cond.Broadcast()
				cond.L.Unlock()
			}
		}(i)
	}
	select {}
}
