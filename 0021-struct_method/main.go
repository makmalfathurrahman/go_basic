package main

import "fmt"

func main() {
	data := Customer{"Alex", "alex@email.com", 10, []string{"Read", "Watch", "Cook"}}
	data.sayHello("James")
}

type Customer struct {
	Name, Email string
	Age         int
	Hobbies     []string
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, ", my name is", customer.Name)
	fmt.Println("My email is", customer.Email)
	fmt.Println("I'm", customer.Age, "years old")
	fmt.Println("My hobbies is", customer.Hobbies)
}
