package main

import "fmt"

func main() {
	var name = Name{Name: "Alex"}
	sayHello(name)
}

type HasName interface {
	GetName() string
}

type Name struct {
	Name string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello, my name is", hasName.GetName())
}

func (name Name) GetName() string {
	return name.Name
}
