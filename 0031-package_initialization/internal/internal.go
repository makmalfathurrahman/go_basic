package internal

import "fmt"

var connection string = "INTERNAL CONNECTION"

func init() {
	fmt.Println(connection)
}
