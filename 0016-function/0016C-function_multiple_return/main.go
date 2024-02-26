package main

import "fmt"

func main() {
	number, integer := call()
	fmt.Println(number)
	fmt.Println(integer)

	fmt.Println("----------")

	_, hello2 := hello()
	fmt.Println(hello2)
}

func call() (string, int) {
	return "one", 1
}

func hello() (string, string) {
	return "Hello", "World!"
}
