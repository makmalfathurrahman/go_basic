package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	runApp(false)
}

func endApp() {
	fmt.Println("-Finish Code-")
	var message = recover()
	fmt.Println("ERROR:", message)
}

func runApp(isTrue bool) {
	defer endApp()

	if !isTrue {
		panic("Code ERROR!")
	}
}
