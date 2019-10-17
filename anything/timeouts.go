package anything

import (
	"fmt"
	"time"
)

/**
* @Author: 胡大海
* @Date: 2019-10-16 21:30
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

// timeout在连接外部资源或者需要限制程序的执行时间中是非常重要的
// 由于存在channel和select，所以在go中实现timeouts是非常简单和优雅的

// 原文中使用了带有缓存的channel
// 说是可以不阻塞第一个goroutine的执行
// 说是这是一个常用的用于避免channel永远不被读取的时候而导致goroutine泄露的模式

// 说白了就是让程序启动的所有goroutine全部顺利执行完成，避免成为孤儿进程

// 找时间测试一下，到底有没有可能成为孤儿进程以及如何检测 todo，进行扩展
func TimeoutsExp() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "work done"
	}()
	select {
	case msg := <-c1:
		fmt.Println("receive msg is ", msg)
	case <- time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}
