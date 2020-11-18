package main

import (
	"fmt"
	"reflect"
	"testing"
)

//type Student struct {
//	Name []string `form:"name"`
//	Age  int64     `form:"age"`
//}


func TestStruct(t *testing.T) {
	//var user *Student
	//value := reflect.ValueOf(user)
	//aaa := reflect.ValueOf(&user)
	//ptr := value
	//ptr = reflect.New(value.Type().Elem())
	//userValue := ptr.Elem()
	//userValue.Set(reflect.ValueOf(Student{[]string{"zhangsan", "lisi"}, 10}))
	//aaa.Elem().Set(ptr)
	//fmt.Println("userValue is ", userValue)
	//fmt.Println("user is ", user)
}

func TestStructVar(t *testing.T) {
	student := Student{}
	val := reflect.ValueOf(&student).Elem()
	val.Field(0).Set(reflect.ValueOf([]string{"zhangsan", "lisi"}))
	fmt.Println("kind is ", val.Field(0).Kind())
	agePtr := reflect.New(val.Field(1).Type().Elem())
	agePtr.Elem().SetInt(100)
	val.Field(1).Set(agePtr)
	fmt.Println("updated student is ", student)
	fmt.Println("updated student's age is ", *student.Age)
}
