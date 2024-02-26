package main

import "fmt"

func main() {
	total1 := sum(10, 10, 10, 10, 10)
	fmt.Println(total1)

	numSlice := []int{10, 10, 10}
	total2 := sum(numSlice...)
	fmt.Println(total2)
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
