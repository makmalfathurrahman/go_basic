package main

import "fmt"

func main() {
	defer fmt.Println("DONE")

	fmt.Println("one")
	fmt.Println("two")
	fmt.Println("three")
}
