package anything

import (
	"context"
	"fmt"
	"time"
)

/**
* @Author: 胡大海
* @Date: 2019-10-17 21:44
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

func AnomalyDetectionExp() {
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <- stop:
				fmt.Println("监控退出，停止了。。。")
				return
			default:
				fmt.Println("goroutine监控中")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("可以通知监控停止")
	stop <- true
	time.Sleep(5 * time.Second)
}

func ContextExp() {
	// context.Background() 返回一个空的Context
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <- ctx.Done():
				fmt.Println("监控退出，停止了....")
				return
			default:
				fmt.Println("goroutine监控中")
				time.Sleep(time.Second * 1)
			}
		}
	}(ctx)
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	time.Sleep(3 * time.Second)
}


func ContextWithManyExp() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "【监控1】")
	go watch(ctx, "【监控2】")
	go watch(ctx, "【监控3】")
	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <- ctx.Done():
			fmt.Println(name, "监控退出，停止了。。。")
			return
		default:
			fmt.Println(name, "goroutine 监控中")
			time.Sleep(2 * time.Second)
		}
	}
}
