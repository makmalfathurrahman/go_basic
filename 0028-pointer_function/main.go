package main

import "fmt"

func main() {
	var address = Address{}
	changeAddres(&address)

	fmt.Println(address)
}

type Address struct {
	Country string
	Planet  string
}

func changeAddres(address *Address) {
	address.Country = "Spain"
	address.Planet = "Saturn"
}
