package designPatterns

import "fmt"

/**
* @Author: 胡大海
* @Date: 2020-10-05 11:32
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

type Product interface {
	create()
}

type Product1 struct {
}

func (p Product1) create() {
	fmt.Println("create Product1")
}

type Product2 struct {
}

func (p Product2) create() {
	fmt.Println("create Product2")
}

func factory(name string) Product {
	switch name {
	case "product1":
		return Product1{}
	case "product2":
		return Product2{}
	default:
		panic("error product")
	}
}
