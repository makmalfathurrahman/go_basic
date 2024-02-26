package main

import "fmt"

func main() {
	var number32 int32 = 32768
	var number16 int16 = int16(number32)

	var name string = "Hello"
	var char byte = name[0]
	var charString string = string(char)

	fmt.Println(number32)
	fmt.Println(number16)

	fmt.Println(name)
	fmt.Println(char)
	fmt.Println(charString)
}
