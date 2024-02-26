package main

import "fmt"

func main() {
	fmt.Println(call(1))
}

func call(number int) bool {
	if number == 1 {
		return true
	} else {
		return false
	}
}
