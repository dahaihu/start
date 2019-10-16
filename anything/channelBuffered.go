package anything

import "fmt"

/**
* @Author: 胡大海
* @Date: 2019-10-11 09:18
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

func ChannelBufferedExp() {
	messages := make(chan string, 2)
	messages <- "hushichang"
	messages <- "limanman"
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
