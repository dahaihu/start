package anything

/**
* @Author: 胡大海
* @Date: 2019-07-28 11:34
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) ReflectCallFunc() {
	fmt.Println("Allen.Wu ReflectCallFunc")
}

func TestReflect() {

	user := User{1, "Allen.Wu", 25}
	// 地址？，所以需要ele？
	sus := reflect.TypeOf(&User{})
	elem := sus.Elem()
	for i:=0; i<elem.NumField(); i++ {
		fmt.Println(elem.Field(i).Name)
	}
	inf := &User{2, "zhangsan", 30}
	value := reflect.ValueOf(inf).Elem()
	// 由于inf是指针类型，所以需要调用Elem()才能获取其真身
	tmp := reflect.TypeOf(inf).Elem()
	// 类型会返回两个值
	field, ok := tmp.FieldByName("Name")
	// 而值只会返回一个值，在不存在的时候会
	_value := value.FieldByName("Name")
	fmt.Println(field, ok)
	fmt.Println(field.Name)
	fmt.Println(_value)
	fmt.Printf("tmp is %v\n", tmp)
	fmt.Printf("变量个数 is %v\n", tmp.NumField())
	fmt.Printf("方法个数 is %v\n", tmp.NumMethod())
	fmt.Println("*****************************************8")
	fmt.Println(_value)
	fmt.Printf("value is %v\n", value.Interface())
	fmt.Printf("值得个数 is %v\n", value.NumField())
	fmt.Printf("方法个数 is %v\n", value.NumMethod())

	for i:=0; i<tmp.NumField(); i++{
		fmt.Printf("type is %v, name is %v, value is %v\n", tmp.Field(i).Type, tmp.Field(i).Name, value.Field(i).Interface())
	}
	// 值？，所以直接调用？
	_type := reflect.TypeOf(user)
	for i:=0; i<_type.NumField(); i++ {
		fmt.Println(_type.Field(i).Name, _type.Field(i).Type)
	}
	DoFiledAndMethod(user)

}

// 只有interface才有反射一说，因为非接口就是已知类型
func DoFiledAndMethod(input interface{}) {

	getType := reflect.TypeOf(input)
	fmt.Println("get Type is :", getType.Name())
	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue)

	// 获取方法字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < getType.NumField(); i++ {
		// 通过type中获取字段的名字和类型
		// 返回的是 StructField 类型
		field := getType.Field(i)
		// 通过value来获取对应字段的值
		// 返回的是 Value 类型
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
