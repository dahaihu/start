package anything

import "fmt"

/**
* @Author: 胡大海
* @Date: 2019-10-11 08:05
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

// 函数中channel是连接并发执行的goroutines的管道
// 你可以在一个goroutine中给channel发送值，然后在其他的goroutine中接受此channel中的值

func ChannelExp() {
	messages := make(chan string)
	go func() { messages <- "ping" }()
	msg := <-messages
	fmt.Println(msg)
}
