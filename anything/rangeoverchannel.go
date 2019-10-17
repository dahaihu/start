package anything

import "fmt"

/**
* @Author: 胡大海
* @Date: 2019-10-17 09:47
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

func RangeOverChannelExp() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	for ele := range queue {
		fmt.Println(ele)
	}
}
