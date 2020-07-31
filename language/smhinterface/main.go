package main

import (
	"fmt"
)

type InterfaceA interface {
	A()
}

type InterfaceB interface {
	B()
}

type Person struct {
	Name string
}

func (p *Person) A() {
	fmt.Println("A")
}

func main() {
	person := &Person{Name: "周二"}
	var a InterfaceA = person
	a.A()
	// var b InterfaceB = a  // 这条语句报错
	// b.B()
}
