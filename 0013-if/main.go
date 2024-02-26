package main

import "fmt"

func main() {
	slice := []int{
		1,
		2,
		3,
		4,
		5,
	}

	if length := len(slice); length == 10 {
		fmt.Println(len(slice) > 10)
	} else if length == 0 {
		fmt.Println(len(slice), "== 0")
	} else {
		fmt.Println(len(slice))
	}
}
