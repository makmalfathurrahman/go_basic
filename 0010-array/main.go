package main

import "fmt"

func main() {
	var array = [3]string{
		"zero",
		"one",
		"two",
	}

	fmt.Println(array)
	fmt.Println(array[0])
	fmt.Println(array[1])
	fmt.Println(array[2])
	fmt.Println(len(array))

	array[2] = "TWO"
	fmt.Println(array)
	fmt.Println(array[2])
}
