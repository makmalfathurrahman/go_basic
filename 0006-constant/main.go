package main

import "fmt"

func main() {
	const lion string = "Lion"
	// lion = "Little Lion" // Disabled

	const (
		name  string = "Alex"
		score int    = 100
	)
	// name = "Brandon" // Disabled
	// score = 50       // Disabled

	fmt.Println(lion)

	fmt.Println(name)
	fmt.Println(score)
}
