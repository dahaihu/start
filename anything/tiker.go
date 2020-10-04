package anything

import (
	"fmt"
	"time"
)

/**
* @Author: 胡大海
* @Date: 2020-02-13 08:46
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */


func TickerStudy() {
	ticker := time.NewTicker(time.Millisecond * 500)
	done := make(chan bool)
	go func() {
		//// 为什么走不下去？？？
		//for t := range ticker.C {
		//	fmt.Println("ticker tick at ", t)
		//}
		//fmt.Println("ticker stopped ???")
		//<- done
		for {
			select {
			case <- done:
				break
			case t := <- ticker.C:
				fmt.Println("tick at ", t)
			}
		}
	}()
	time.Sleep(time.Second * 3)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
