package anything

import (
	"fmt"
	"time"
)

/**
* @Author: 胡大海
* @Date: 2019-10-11 09:24
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

func worker(work chan bool) {
	fmt.Println("start working")
	time.Sleep(time.Second)
	fmt.Println("work done")
	work <- true
}

func ChannelSynchronizationExp() {
	work := make(chan bool)
	go worker(work)
	// 如果不执行下面这句，那么worker甚至还没启动，程序就已经结束了
	<- work
}
