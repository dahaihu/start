package anything

import (
	"fmt"
	"unsafe"
)

/**
* @Author: 胡大海
* @Date: 2019-10-12 15:18
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

func MapExp() {
	// 传递变量的指针v
	i := 10
	modifyInterger(&i)
	fmt.Println(i)
	persons := make(map[string]int)
	persons["胡世昌"] = 10
	persons["李漫漫"] = 11
	fmt.Println("persons is ", persons)
	_, ok := persons["nihao"]
	fmt.Println("nihao exists in persons ? ", ok)
	modifyMap(persons)
	fmt.Println("size is ", unsafe.Sizeof(persons))
	fmt.Println("update in function's result is ", persons)

	dict := map[string]int{"胡世昌": 100, "李漫漫": 1000}
	fmt.Println("dict is ", dict)
	fmt.Println("new and map ", &dict)
	array := []int{1, 2, 3, 5}
	fmt.Println(unsafe.Sizeof(array))
	modifyArray(array)
}

func modifyInterger(i *int) {
	*i = 11
}

func modifyMap(dict map[string]int) {
	fmt.Println("size is ", unsafe.Sizeof(dict))
	dict["胡大海"] = 100
}

func modifyArray(array []int) {
	fmt.Println("array is ", unsafe.Sizeof(array))
}

