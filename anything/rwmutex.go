package anything

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

/**
* @Author: 胡大海
* @Date: 2019-10-26 09:52
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

// 代码地址：https://medium.com/golangspec/sync-rwmutex-ca6c6c3208a0

// 如何表示内部的状态呢？
// RLock 和 Lock 锁住的到底是什么呢？
// 还是开发人员自己定义这个锁住的东西？
// 如何定义一轮呢？？

func init() {
	rand.Seed(time.Now().Unix())
}

func sleep() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
}

func reader(wg *sync.WaitGroup, rw *sync.RWMutex, readerChan chan int) {
	rw.RLock()
	readerChan <- 1
	sleep()
	readerChan <- -1
	rw.RUnlock()
	wg.Done()
}

func writer(wg *sync.WaitGroup, rw *sync.RWMutex, writerChan chan int) {
	rw.Lock()
	writerChan <- 1
	sleep()
	writerChan <- -1
	rw.Unlock()
	wg.Done()
}

func RWLockExp() {
	readerChan := make(chan int)
	writerChan := make(chan int)
	wg := sync.WaitGroup{}
	rw := sync.RWMutex{}
	var readerCount, writerCount int
	go func() {
		for {
			select {
			case val := <-readerChan:
				readerCount += val
			case val := <-writerChan:
				writerCount += val
			}
			fmt.Printf("%s%s\n", strings.Repeat("R", readerCount), strings.Repeat("W", writerCount))
		}
	}()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go writer(&wg, &rw, writerChan)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go reader(&wg, &rw, readerChan)
	}

	wg.Wait()

}
