package main

import (
	"bytes"
	"fmt"
	"math"
	"os/exec"
)

func main() {
	var bts bytes.Buffer
	cmd := exec.Command("ls", "-alht")
	cmd.Stdout = &bts
	cmd.Run()
	fmt.Print(bts.String())
	var a int64 = math.MaxInt64
	fmt.Printf("a is %d\n", a)
}
