package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := Sum(1, 0)
	if err == nil {
		fmt.Println("Result =", result)
	} else {
		fmt.Println("ERROR:", err.Error())
	}
}

func Sum(num1, num2 int) (int, error) {
	if num1 == 0 || num2 == 0 {
		return 0, errors.New("number cannot be 0")
	} else {
		sum := num1 + num2
		return sum, nil
	}
}
