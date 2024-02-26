package main

import "fmt"

func main() {
	runApp(true)
}

func endApp() {
	fmt.Println("-Finish Code-")
}

func runApp(isTrue bool) {
	defer endApp()

	if !isTrue {
		panic("Code ERROR!")
	}
}
