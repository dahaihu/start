package main

import "fmt"

type mac struct{}

func (m *mac) insertInSquarePort() {
	fmt.Println("mac insert in square port")
}

var _ computer = &mac{}
