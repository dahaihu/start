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
		//for ele := range jobs {
		//	fmt.Println("receive job ", ele)
		//}
		//done <- true
		for {
			// 感觉换成ok更合适，ok为true的时候，表示有值出来，
			// 否则jobs就被close了，并且为空
			j, ok := <-jobs
			fmt.Println("status is ", j, ok)
			if ok {
				fmt.Println("received job ", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				break
			}
		}
	}()
	for i:=0; i<3; i++ {
		jobs <- i+1
		fmt.Println("sent job ", i+1)
	}
	//close(jobs)
	fmt.Println("sent all jobs")
	// 由于是没有长度的channel，所以接受的时候，如果channel中没有值，就会一直在这阻塞
	// 直到channel中准备传递值为止， 这样也就是相当于完成了等待goroutine执行完成的过程
	<- done
}
