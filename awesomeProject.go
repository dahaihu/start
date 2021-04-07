package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	s := uuid.New().String()
	fmt.Printf("share code is: %s\n", s)
	// 表格定义的时候，去掉字符集的部分
}
