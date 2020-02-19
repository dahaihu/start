package anything

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/**
* @Author: 胡大海
* @Date: 2020-02-18 08:24
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

func MutexExp() {
	var readOps uint64
	var writeOps uint64
	// 不是在传入给函数使用的情况下，用地址还是用值都是一样的
	mutex := sync.Mutex{}
	state := make(map[int]int)

	for i := 0; i < 100; i++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				key := rand.Intn(5)
				value := rand.Intn(100)
				mutex.Lock()
				state[key] = value
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("read ops is ", atomic.LoadUint64(&readOps))
	fmt.Println("read ops is ", atomic.LoadUint64(&writeOps))
	mutex.Lock()
	fmt.Println("state is ", state)
	mutex.Unlock()

}
