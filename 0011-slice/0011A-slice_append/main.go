package main

import "fmt"

func main() {
	var days = [...]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	var slice1 = days[5:]
	slice1[0] = "NEW Friday"
	slice1[1] = "NEW Saturday"
	fmt.Println(slice1)

	var slice2 = append(slice1, "HOLIDAY")
	fmt.Println(slice2)
}
