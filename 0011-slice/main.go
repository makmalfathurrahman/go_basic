package main

import "fmt"

func main() {
	months := [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	slice1 := months[4:6]
	fmt.Println(slice1)
	fmt.Println("len =", len(slice1))
	fmt.Println("cap =", cap(slice1))
	fmt.Println(slice1[0])
	fmt.Println(slice1[1])

	fmt.Println("----------")

	slice2 := months[6:12]
	fmt.Println(slice2)
	fmt.Println("len =", len(slice2))
	fmt.Println("cap =", cap(slice2))
	fmt.Println(slice2[0])
	fmt.Println(slice2[1])
	fmt.Println(slice2[2])

	fmt.Println("----------")

	slice3 := months[:2]
	fmt.Println(slice3)
	fmt.Println("len =", len(slice3))
	fmt.Println("cap =", cap(slice3))

	fmt.Println("----------")

	var slice4 []string = months[4:]
	fmt.Println(slice4)
	fmt.Println("len =", len(slice4))
	fmt.Println("cap =", cap(slice4))

	fmt.Println("----------")

	var slice5 []string = months[:]
	fmt.Println(slice5)
	fmt.Println("len =", len(slice5))
	fmt.Println("cap =", cap(slice5))

}
