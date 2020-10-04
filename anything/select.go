package anything

import (
	"fmt"
	"time"
)

/**
* @Author: 胡大海
* @Date: 2019-10-16 21:17
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */


// go中的select是的你可以等待多个channel的操作
// 结合goroutine和channel以及select是go非常强大的一个特性


// 我们可以使用select同时等待c1、c2两个channel的值
func SelectExp() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second)
		c1 <- "msg1"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "msg2"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case msg1 := <- c1:
			fmt.Println("receive msg is ", msg1)
		case msg2 := <- c2:
			fmt.Println("receive msg is ", msg2)
		default:
			fmt.Println("no channel is ready")
		}
	}
}
