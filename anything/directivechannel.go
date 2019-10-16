package anything

import "fmt"

/**
* @Author: 胡大海
* @Date: 2019-10-16 21:03
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

// 可以使用channels作为函数的参数，你甚至可以指定channel是用来发送或者接受值的
// 通过这种方式可以增加程序的安全性

// ping函数的参数pings chan<- string，表示的是pings是仅仅用来接受值的，可以用参数代替chan的位置来思考参数的方向
// 例如下面的替换结果就是 pings <- string, 那么显然这个pings就是用来接受值的
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 下面的pongs参数替换的结果就是 <-pings，那么pings就是用来发送的了
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}



// 也可以理解参数中chan箭头的朝向吧，朝向chan就是仅仅用于接受，朝离chan就是仅仅用于发送
func DirectiveChannelExp() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "胡世昌")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
