package main

import "fmt"

func main() {
	// 1st Method
	var data Customer
	data.Name = "Alex"
	data.Email = "alex@email.com"
	data.Age = 10
	data.Hobbies = []string{"Read"}

	fmt.Println(data)
	fmt.Println(data.Name)
	fmt.Println(data.Email)
	fmt.Println(data.Age)

	for i := 0; i < len(data.Hobbies); i++ {
		item := data.Hobbies[i]
		fmt.Println("-", item)
	}

	fmt.Println("----------")

	// 2nd Method
	data = Customer{
		Name:    "Brandon",
		Email:   "brandon@email.com",
		Age:     20,
		Hobbies: []string{"Read", "Watch"},
	}
	fmt.Println(data)
	fmt.Println(data.Name)
	fmt.Println(data.Email)
	fmt.Println(data.Age)

	for i := 0; i < len(data.Hobbies); i++ {
		item := data.Hobbies[i]
		fmt.Println("-", item)
	}

	fmt.Println("----------")

	// 3rd Method
	data = Customer{"Chris", "chris@email.com", 30, []string{"Read", "Watch", "Cook"}}
	fmt.Println(data)
	fmt.Println(data.Name)
	fmt.Println(data.Email)
	fmt.Println(data.Age)

	for i := 0; i < len(data.Hobbies); i++ {
		item := data.Hobbies[i]
		fmt.Println("-", item)
	}
}

type Customer struct {
	Name, Email string
	Age         int
	Hobbies     []string
}
