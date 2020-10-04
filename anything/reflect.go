package anything

/**
* @Author: 胡大海
* @Date: 2019-07-28 11:34
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

import (
	"fmt"
	"reflect"
	"runtime"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type Person struct {
	Id   int
	Name string
	Age  int
}


func (u User) ReflectCallFunc() {
	fmt.Println("Allen.Wu ReflectCallFunc")
}

func TestConvert() {
	pc, file, line, ok := runtime.Caller(1)
	fmt.Printf("%v, %v, %v, %v\n", pc, file, line, ok)
	// 强制类型转换就相当于新创建一个对象，因为两者的地址不同
	user := User{1, "Allen.Wu", 25}
	user1 := Person(user)
	user2 := Person(user)
	fmt.Printf("%p, %p\n", &user1, &user2)
}

func TestReflect() {
	user := User{10, "hudahai", 10}
	DoFiledAndMethod(user)
}

// 只有interface才有反射一说，因为非接口就是已知类型
func DoFiledAndMethod(input interface{}) {
	getType := reflect.TypeOf(input)
	fmt.Printf("type is %v\n", getType)
	getValue := reflect.ValueOf(input)
	fmt.Printf("type is %v\n", getValue)

	for i:=0; i<getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i)
		fmt.Printf("name is %v, type is %v, value is %v\n", field.Name, field.Type, value)
		fmt.Printf("interface's effect is %v\n", value == value.Interface())
		fmt.Printf("value is %T, interface value is %T\n", value, value.Interface())
	}

	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
