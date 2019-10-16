package anything

import (
	"fmt"
	"time"
)

/**
* @Author: 胡大海
* @Date: 2019-07-29 16:51
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

func TestTime() {
	timeUnix := time.Now().Unix()
	formatTimeStr := time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05")
	fmt.Println(formatTimeStr) //打印结果：2017-04-11 13:30:39
}
