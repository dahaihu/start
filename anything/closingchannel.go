package anything

import "fmt"

/**
* @Author: 胡大海
* @Date: 2019-10-17 09:40
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

func ClosingChannelExp() {
	jobs := make(chan int, 5)
	done := make(chan bool)
	go func() {
		for {
			job, ok := <-jobs
			if ok {
				fmt.Println("received job ", job)
			} else {
				fmt.Println("jobs channel closed")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		jobs <- i
		fmt.Println("send job ", i)
	}
	close(jobs)
	<-done
}

func closingChannelSend() {
	mark := make(chan int, 5)
	for i := 0; i < 5; i++ {
		mark <- i
	}
	close(mark)
	var res int
	//var ok bool
	for res = range mark {
		fmt.Println("res is ", res)
	}
}
