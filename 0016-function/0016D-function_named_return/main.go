package main

import "fmt"

func main() {
	result := applyFunction(3, 4, add)
	fmt.Println("Result =", result)

	result = applyFunction(7, 2, func(a, b int) int {
		return a - b
	})
	fmt.Println("Result =", result)
}

type Function func(int, int) int

func applyFunction(a, b int, function Function) int {
	return function(a, b)
}

func add(a, b int) int {
	return a + b
}
