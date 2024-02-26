package main

import "fmt"

func main() {
	var varString string = "This is String"
	var varNumber int = 10
	var varBoolean bool = true

	var (
		name  string = "Alex"
		score int    = 100
	)

	fmt.Println(varString)
	fmt.Println(varNumber)
	fmt.Println(varBoolean)

	fmt.Println(name)
	fmt.Println(score)
}
