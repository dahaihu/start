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
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}

func reader(c chan int, m *sync.RWMutex, wg *sync.WaitGroup) {
	defer wg.Wait()
	m.RLock()
	c <- 1
	sleep()
	c <- -1
	m.RUnlock()
}

func writer(c chan int, m *sync.RWMutex, wg *sync.WaitGroup) {
	defer wg.Wait()
	m.Lock()
	c <- 1
	sleep()
	c <- -1
	m.Unlock()
}

func RWLockExp() {
	rw := sync.RWMutex{}
	readerChan := make(chan int)
	writerChan := make(chan int)
	var readerCount, writerCount int
	go func() {
		for {
			select{
			case val := <- readerChan:
				readerCount += val
			case val := <- writerChan:
				writerCount += val
			}
			fmt.Printf("%s%s\n", strings.Repeat("W", writerCount), strings.Repeat("R", readerCount))
		}
	}()

	wg := sync.WaitGroup{}

	for i:=0; i<3; i++ {
		wg.Add(1)
		go writer(writerChan, &rw, &wg)
	}

	for i:=0; i<10; i++ {
		wg.Add(1)
		go reader(readerChan, &rw, &wg)
	}

	wg.Wait()

}
