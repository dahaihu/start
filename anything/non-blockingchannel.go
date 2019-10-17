package anything

import (
	"fmt"
)

/**
* @Author: 胡大海
* @Date: 2019-10-16 21:41
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

func NoneBlockingChannelOperationExp() {
	messages := make(chan string)
	signals := make(chan bool)
	// 第一次，由于messages是个没有缓存的channel，但是呢也没有值，所以是default执行
	select {
	case msg := <-messages:
		fmt.Println("received message is ", msg)
	default:
		fmt.Println("no message received")
	}
	// select中case的含义到底是什么呢？
	// 是顺序查看每个case的表达式可以立刻的执行吗
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message ", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signale", sig)
	default:
		fmt.Println("no activity")
	}
}
