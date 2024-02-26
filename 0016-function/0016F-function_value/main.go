package main

import "fmt"

func main() {
	number, integer, boolean := name()
	fmt.Println(number)
	fmt.Println(integer)
	fmt.Println(boolean)
}

func name() (number string, integer int, boolean bool) {
	number = "one"
	integer = 1
	boolean = true

	return number, integer, boolean
}
