package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
	"log"
)

func lock() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()

	s2, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s2.Close()

	m1 := concurrency.NewMutex(s1, "/my-lock/")
	m2 := concurrency.NewMutex(s2, "/my-lock/")
	if err := m1.Lock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("acquire lock for m1")
	m2Locked := make(chan struct{})
	go func() {
		defer close(m2Locked)
		if err := m2.Lock(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	if err := m1.Unlock(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("release lock for m1")
	<-m2Locked
	fmt.Println("acquire lock for m2")
}
