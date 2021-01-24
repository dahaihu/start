package main

import "fmt"

func safeRun(f func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err is ", err)
		}
	}()
	f()
}
func main() {
	safeRun(func() {
		panic("wo ri ni da ye ")
	})
	fmt.Println("safe end")
}
