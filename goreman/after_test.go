package main

import (
	"fmt"
	"testing"
	"time"
)

func TestAfter(t *testing.T) {
	timer := time.AfterFunc(time.Second, func() {
		fmt.Println("after doing")
	})
	timer.Stop()
	fmt.Println("stop timer")
}
