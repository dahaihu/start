package anything

import (
	"fmt"
	"unsafe"
)

/**
* @Author: 胡大海
* @Date: 2020-04-05 16:28
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

//
// runtime/slice.go
type slice struct {
	array unsafe.Pointer // 元素指针
	len   int            // 长度
	cap   int            // 容量
}

func UnSafeExp() {
	s := "hushichang"
	fmt.Println(unsafe.Pointer(&s))


	n := Num{"test", 10}

	x := (*[2]uintptr)(unsafe.Pointer(&n))
	fmt.Println(*(*string)(unsafe.Pointer((*[2]uintptr)(unsafe.Pointer(&n))[0])))
	h := [2]uintptr{x[1], x[0]}
	fmt.Println("aaa", (*ReverseNum)(unsafe.Pointer(&h)))


	nPointer := unsafe.Pointer(&n)
	niPointer := (*string)(unsafe.Pointer(uintptr(nPointer) + unsafe.Offsetof(n.name)))
	*niPointer = "煎鱼"
	njPointer := (*int64)(unsafe.Pointer(uintptr(nPointer) + unsafe.Offsetof(n.value)))
	*njPointer = 100
	fmt.Println("n is ", n)
}

type Num struct {
	name  string
	value int64
}

type ReverseNum struct {
	value int64
	name string
}


