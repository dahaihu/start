package anything

import (
	"fmt"
	"math/rand"
	"time"
)

/**
* @Author: 胡大海
* @Date: 2020-02-29 11:28
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

func BasicRateLimiting() {
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(200 * time.Millisecond)
	for request := range requests {
		<-limiter
		fmt.Println("Process request", request, time.Now())
	}
}

func BurstRateLimiting() {
	// 所谓的令牌桶算法
	// 以固定的速率 往一个桶子里面传递数据
	// 来请求，在令牌桶有令牌的情况下直接操作
	// 令牌桶没有数据的话，就会直到有令牌为止
	burstLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstLimiter <- t
		}
	}()

	burstRequests := make(chan int, 100)
	done := make(chan bool)
	go func() {
		for request := range burstRequests {
			<- burstLimiter
			fmt.Println("processing ", request, time.Now())
		}
		done <- true
	}()

	for i := 0; i < 100; i++ {
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		burstRequests <- i
	}
	close(burstRequests)
	<- done
}
