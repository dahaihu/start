package anything

import "fmt"

/**
* @Author: 胡大海
* @Date: 2019-10-09 21:31
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

type person struct {
	name string
	age int
}

// 结构体是如何初始化的，和python有点相似，一个用小括号，一个用大括号
func NewPerson(name string) *person {
	p := person{name, 32}
	p.age = 42
	return &p

	// 还可以这个样子写
	// // return &person{name, 42}
}

func StructExp() {
	p := NewPerson("胡世昌")
	fmt.Println("person is ", p)
	// 使用new来初始化struct，返回的是初始化该实例的一个指针
	p = new(person)
	p.name = "胡大海"
	p.age = 10
	fmt.Println(p)
}
