package main

import "fmt"

func main() {
	var hello = Hello{"World!"}
	hello.getHello()

	fmt.Println(hello.Hello)
}

type Hello struct {
	Hello string
}

func (hello *Hello) getHello() {
	hello.Hello = "Hello " + hello.Hello
}
