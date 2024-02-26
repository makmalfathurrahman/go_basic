package main

import "fmt"

func main() {
	hello()
	count(1, 2)
	fmt.Println(call(1))
}

func hello() {
	fmt.Println("Hello World!")
}

func count(a int, b int) {
	fmt.Println(a + b)
}

func call(number int) bool {
	if number == 1 {
		return true
	} else {
		return false
	}
}
