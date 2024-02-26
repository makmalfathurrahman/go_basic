package main

import "fmt"

func main() {
	for i := 0; i <= 50; i += 10 {
		fmt.Println(i)
	}

	fmt.Println("----------")

	for i := 0; i <= 50; i += 10 {
		if i == 30 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("----------")

	for i := 0; i <= 50; i += 10 {
		if i == 30 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("----------")

	number := []string{"zero", "one", "two", "three", "four"}
	for index, val := range number {
		if val == "two" {
			continue
		}
		fmt.Printf("%v\t%v\n", index, val)
	}

	fmt.Println("----------")

	integer := []string{"zero", "one", "two", "three", "four"}
	for _, val := range integer {
		fmt.Println(val)
	}
}
