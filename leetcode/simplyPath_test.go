package leetcode

import "fmt"

func ExampleSimplifyPath() {
	fmt.Println(simplifyPath("/a//b////c/d//././/.."))
	// Output: /a/b/c
}