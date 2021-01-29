package main

import (
	"flag"
	"fmt"
)

var name = flag.String("f", "zhangsan", "your name")

var age = flag.Uint("a", 100, "your age")

func tag() {
	flag.Parse()
	fmt.Printf("command line args are %v\n", flag.Args())
	fmt.Printf("name is %v, age is %d\n", *name, *age)
}
