package designPattern

import "fmt"

type Factory struct {
}

type Product interface {
	create()
}

type Product1 struct {
}

func (p1 Product1) create() {
	fmt.Println("this is product1")
}

type Product2 struct {
}

func (p2 Product2) create() {
	fmt.Println("this is product2")
}

type Product3 struct {
}

func (p3 Product3) create() {
	fmt.Println("this is product3")
}

func (f Factory) Generate(name string) Product {
	switch name {
	case "product1":
		return Product1{}
	case "product2":
		return Product2{}
	case "product3":
		return Product3{}
	default:
		return nil
	}
}

func factory() {
	f := new(Factory)
	p1 := f.Generate("product1")
	p1.create()
	p2 := f.Generate("product2")
	p2.create()
}
