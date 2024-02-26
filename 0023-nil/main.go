package main

import "fmt"

func main() {
	fmt.Println(Map("Hello"))

	slice := []int{1, 2, 3}
	fmt.Println(Slice(slice))
}

func Map(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func Slice(number []int) []int {
	if number == nil {
		return nil
	} else {
		var newSlice []int = make([]int, 0, len(number)+1)
		appendSlice := append(newSlice, number...)
		fmt.Println(len(number) + 1)
		return appendSlice
	}
}
