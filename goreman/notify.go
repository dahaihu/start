package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal, 10)
	// 可以在命令行进行 ctr+c 的操作捕捉信号
	signal.Notify(c, syscall.SIGINT)

	for s := range c {
		fmt.Println("received", s)
	}
}
