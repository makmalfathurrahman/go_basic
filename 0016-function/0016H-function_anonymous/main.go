package main

import "fmt"

func main() {
	var greet = func() {
		fmt.Println("Hello World!")
	}
	greet()

	var sum = func(n1, n2 int) {
		var sum int = n1 + n2
		fmt.Println("Sum =", sum)
	}
	sum(1, 2)

	var sumReturn = func(n1, n2 int) int {
		var sumReturn int = n1 + n2
		return sumReturn
	}
	fmt.Println("Sum Return =", sumReturn(1, 3))
}
