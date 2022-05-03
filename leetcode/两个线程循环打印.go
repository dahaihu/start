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
		for i := 0; i < times; i++ {
			<-two
			fmt.Println(2)
			if i == times-1 {
				close(one)
				close(two)
				break
			}
			one <- signal
		}
		wg.Done()
	}()
	one <- signal
	wg.Wait()
}

func twoThreadPrintUsingSignal() {
	var num int
	cond := sync.NewCond(new(sync.Mutex))
	count := 2
	for i := 0; i < count; i++ {
		go func(idx int) {
			for {
				cond.L.Lock()
				for  {
					fmt.Printf("nocondition: idx %d get num %d\n", idx, num)
					if num == idx {
						break
					}
					cond.Wait()
				}
				fmt.Printf("condition: idx %d get %d\n", idx, num)
				if num = num + 1; num == count {
					num = 0
				}
				cond.Broadcast()
				cond.L.Unlock()
			}
		}(i)
	}
	select {}
}
