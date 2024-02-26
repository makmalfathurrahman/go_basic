package main

import "fmt"

func main() {
	operation(true, false)
	fmt.Println("----------")

}

func operation(a bool, b bool) {
	fmt.Println("-----AND &&-----")
	fmt.Println(a && a)
	fmt.Println(a && b)
	fmt.Println(b && a)
	fmt.Println(b && b)

	fmt.Println("-----OR ||-----")
	fmt.Println(a || a)
	fmt.Println(a || b)
	fmt.Println(b || a)
	fmt.Println(b || b)

	fmt.Println("-----NOT !-----")
	fmt.Println(!a)
	fmt.Println(!b)
}
