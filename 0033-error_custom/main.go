package main

import "fmt"

func main() {
	result := SaveData("ABC", nil)

	if result != nil {
		if validationError, ok := result.(*validationError); ok {
			fmt.Println("Validation error:", validationError.Message)
		} else if notFoundError, ok := result.(*notFoundError); ok {
			fmt.Println("Not found error:", notFoundError.Message)
		} else {
			fmt.Println("Unknown error", result.Error())
		}
	}
}

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (v *notFoundError) Error() string {
	return v.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation error"}
	}

	if id != "ABC" {
		return &notFoundError{"Data not found"}
	}

	return nil
}
