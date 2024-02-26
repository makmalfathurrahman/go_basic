package main

import "fmt"

func main() {
	fmt.Println(factorialLoop(3))
	fmt.Println(factorialRecursice(4))
}

func factorialLoop(num int) int {
	result := 1
	for i := num; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursice(num int) int {
	if num == 1 {
		return 1
	} else {
		return num * factorialRecursice(num-1)
	}
}
