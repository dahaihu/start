package main

import (
	"fmt"
	"reflect"
)

func call(f interface{}, args []interface{}) {
	in := make([]reflect.Value, len(args))
	for i := range args {
		in[i] = reflect.ValueOf(args[i])
	}
	reflect.ValueOf(f).Call(in)
}

func greet(a int64, b string) {
	fmt.Printf("a is %v, b is %v\n", a, b)
}

func main() {
	call(greet, []interface{}{int64(10), "world"})
}
