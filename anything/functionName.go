package anything

/**
* @Author: 胡大海
* @Date: 2019-08-07 13:38
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

import (
	"fmt"
	"log"
	"runtime"
)

func Foo() {
	fmt.Printf("我是 %s, 谁在调用我?\n", printMyName())
	Bar()
}
func Bar() {
	fmt.Printf("我是 %s, 谁又在调用我?\n", printMyName())
	trace()
}
func printMyName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func trace() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	n := runtime.Callers(0, pc)
	for i := 0; i < n; i++ {
		f := runtime.FuncForPC(pc[i])
		file, line := f.FileLine(pc[i])
		fmt.Printf("%s:%d %s\n", file, line, f.Name())
	}
}

func TestCaller() {
	// pc的含义是 program counter
	pc, file, line, ok := runtime.Caller(1)
	// 是否获取成功
	log.Println(ok)
	// 函数指针
	log.Println(pc)
	// 所属文件
	log.Println(file)
	// 所属行
	log.Println(line)
	// 获取函数信息
	f := runtime.FuncForPC(pc)
	// 函数名
	log.Println(f.Name())
}
