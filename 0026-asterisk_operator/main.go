package main

import "fmt"

func main() {
	var address = Address{Country: "Indonesia", Planet: "Earth"}

	addressX := &address
	addressX.Country = "England"

	*addressX = Address{Country: "United States", Planet: "Mars"}
	fmt.Println(addressX)
	fmt.Println(address)
}

type Address struct {
	Country string
	Planet  string
}
