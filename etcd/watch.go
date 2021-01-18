package main

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

func watchKey(cli *clientv3.Client, key string) {
	wch := cli.Watch(cli.Ctx(), key)
	for resp := range wch {
		for idx, ev := range resp.Events {
			fmt.Printf("Event ids is %d, Type: %s Key:%s Value:%s\n",
				idx, ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}

func change() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 3 * time.Second,
	})
	fmt.Println("starting!!!!")
	if err != nil {
		log.Fatal(err)
		return
	}
	key := "/test/b"
	defer cli.Close()
	go watchKey(cli, key)
	for i := 0; i < 10; i++ {
		resp, err := cli.Put(cli.Ctx(),
			"/test/b", "something",
			clientv3.WithPrevKV())
		if err != nil {
			fmt.Printf("put error is %v\n", err)
			continue
		}
		fmt.Printf("put resp is %v\n", resp)
		time.Sleep(time.Microsecond * 100)
	}
	resp, err := cli.Delete(cli.Ctx(), key)
	fmt.Printf("del resp is %v\n", resp)
	time.Sleep(time.Second)
}
