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
	one, two := make(chan struct{}), make(chan struct{})
	signal := struct{}{}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		for {
			_, ok := <-one
			if !ok {
				break
			}
			fmt.Println(1)
			two <- signal
		}
		wg.Done()
	}()
	go func() {
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
