package main

import (
	"fmt"
	"os/exec"
	"sync"
	"syscall"
	"testing"
)
type logger struct {}


func (l logger) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return len(p), nil
}

func TestProcess(t *testing.T) {
	l := logger{}
	cmd := exec.Command("ls", "-lah")
	cmd.Stdout = l
	cmd.Stderr = l
	if err := cmd.Start(); err != nil {
		fmt.Println("start err is ", err)
		return
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		defer wg.Done()
		if err := cmd.Wait(); err != nil {
			fmt.Printf("wait err is %v\n", err)
		}
	}()

	if err := cmd.Process.Signal(syscall.SIGINT); err != nil {
		fmt.Printf("signal err is %v\n", err)
	}
	wg.Wait()
}
