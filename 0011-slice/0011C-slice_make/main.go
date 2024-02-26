package main

import "fmt"

func main() {
	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "zero"
	newSlice[1] = "one"
	fmt.Println(newSlice)
	fmt.Println("len =", len(newSlice))
	fmt.Println("cap =", cap(newSlice))

	fmt.Println("----------")

	appendSlice := append(newSlice, "two", "three", "four")
	fmt.Println(appendSlice)
	fmt.Println(newSlice)
	fmt.Println("len =", len(appendSlice))
	fmt.Println("cap =", cap(appendSlice))

	appendSlice[0] = "ZERO"
	fmt.Println(appendSlice)
	fmt.Println(newSlice)
}
