package main

import (
	"fmt"
	"os/exec"
	"sync"
	"syscall"
	"testing"
	"time"
)

type logger struct{}

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
	go func() {
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

func TestKillProcess(t *testing.T) {
	cmd := exec.Command("/bin/sh", "-c", "sleep 100")
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	start := time.Now()
	time.AfterFunc(10*time.Second, func() {
		_ = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	})
	err := cmd.Run()
	fmt.Printf("pid=%d duration=%s err=%s\n", cmd.Process.Pid, time.Since(start), err)
}
