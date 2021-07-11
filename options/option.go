package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type PersonSetting func(*Person)

func PersonWithName(name string) PersonSetting {
	return func(p *Person) {
		p.Name = name
	}
}

func PersonWithAge(age int) PersonSetting {
	return func(p *Person) {
		p.Age = age
	}
}

func NewPerson(ops ...PersonSetting) *Person {
	person := &Person{}
	for _, op := range ops {
		op(person)
	}
	return person
}

func main() {
	person := NewPerson(PersonWithName("zhangsan"), PersonWithAge(10))
	fmt.Println("client is ", *person)
}
