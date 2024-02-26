package main

import "fmt"

func main() {
	result := random()
	// 1st Method
	resultInt := result.(int)
	fmt.Println(resultInt)

	// 2nd Method
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	case bool:
		fmt.Println("Bool", value)
	default:
		fmt.Println("Unknown")
	}
}

func random() interface{} {
	return 0
}
