package main

import "fmt"

func main() {
	// By Value
	var address = Address{Country: "Indonesia", Planet: "Earth"}
	fmt.Println(address)

	addressX := address
	addressX.Country = "England"
	fmt.Println(addressX)

	// By Reference
	var addressRef Address = Address{Country: "Indonesia", Planet: "Earth"}

	var addressXRef *Address = &addressRef
	addressXRef.Country = "United States"
	fmt.Println(addressXRef)
	fmt.Println(addressRef)
}

type Address struct {
	Country string
	Planet  string
}
