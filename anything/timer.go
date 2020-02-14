package anything

import (
	"fmt"
	"time"
)

/**
* @Author: 胡大海
* @Date: 2020-02-13 08:34
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

// 感觉timer这样用还是不好，会造成goroutine泄露的
// 因为等待timer2的


func TimerTest() {
	timer1 := time.NewTimer(time.Second)
	<-timer1.C
	fmt.Println("timer1 fired")

	timer2 := time.NewTimer(time.Millisecond)
	go func(){
		<- timer2.C
		fmt.Println("timer2 fired")
	}()
	time.Sleep(time.Millisecond)
	// timer为什么返回的是一个bool类型的？
	// 因为在制止timer的时候，timer可能都已经执行完成了
	ok := timer2.Stop()
	if ok {
		fmt.Println("timer2 stopped")
	}
	time.Sleep(time.Second)
}
