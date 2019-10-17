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
			j, more := <-jobs
			if more {
				fmt.Println("received job ", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job ", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	// 由于是没有长度的channel，所以接受的时候，如果channel中没有值，就会一直在这阻塞
	// 直到channel中准备传递值为止， 这样也就是相当于完成了等待goroutine执行完成的过程
	<- done
}
