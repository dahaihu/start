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

func modifyValue() {
	var a float32 = 1.0
	var b float32 = 2.0
	value := reflect.ValueOf(&a)
	point := reflect.TypeOf(a)
	value.Elem().Set(reflect.ValueOf(b))
	fmt.Println(value, point.Kind() == reflect.Float32)
	fmt.Println("updated value is ", a, a == 2.0)
}

func appendIntToSlice(arrPtr interface{}) {
	valuePtr := reflect.ValueOf(arrPtr)
	value := valuePtr.Elem()

	value.Set(reflect.Append(value, reflect.ValueOf(55)))

	fmt.Println(value.Len())
}

func appendToSlice(arrPtr interface{}, value interface{}) error {
	arrValue := reflect.ValueOf(arrPtr).Elem()
	fmt.Println(reflect.TypeOf(value).Kind())
	fmt.Println(reflect.TypeOf(arrPtr).Elem().Elem().Kind())
	arrValue.Set(reflect.Append(arrValue, reflect.ValueOf(value)))
	return nil
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) ReflectCallFunc() {
	fmt.Println("Allen.Wu ReflectCallFunc")
}

// 通过接口来获取任意参数，然后一一揭晓
func DoFiledAndMethod(input interface{}) {

	_type := reflect.TypeOf(input)
	fmt.Println("type is ", _type)

	_value := reflect.ValueOf(input)
	fmt.Println("value is ", _value)

	for idx := 0; idx < _type.NumField(); idx++ {
		fmt.Println(_type.Field(idx).Name, _type.Field(idx).Type,
			_value.Field(idx))
	}
}


func main() {
	//call(greet, []interface{}{int64(10), "world"})
	//modifyValue()
	//a := []int{1, 2, 3}
	//appendIntToSlice(&a)
	//fmt.Println(appendToSlice(&a, 10))
	//fmt.Println(a)
	user := User{1, "Allen.Wu", 25}

	DoFiledAndMethod(user)
}
