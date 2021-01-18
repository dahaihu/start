package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

func transaction() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 3 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed err: ", err)
		return
	}
	defer client.Close()
	key, value := "/transaction", "transaction"
	_, _ = client.Put(client.Ctx(), key, value)

	_, err = client.Txn(context.Background()).
		If(clientv3.Compare(clientv3.Value(key), "=", value)).
		Then(clientv3.OpPut(key, "changed")).
		Commit()
	if err != nil {
		log.Fatal(err)
	}
	if resp, err := client.Get(context.TODO(), key); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}
}
