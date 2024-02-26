package main

import "fmt"

func main() {
	var mapping map[string]int = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	fmt.Println(mapping)
	fmt.Println("len =", len(mapping))
	fmt.Println(mapping["one"])
	fmt.Println(mapping["three"])

	mapping["two"] = 222222
	fmt.Println(mapping)

	fmt.Println("----------")

	number := make(map[int]string)
	number[100] = "One Hundred"
	number[200] = "Two Hundred"
	number[300] = "Three Hundred"

	fmt.Println(number)
	fmt.Println(number[100])

}
