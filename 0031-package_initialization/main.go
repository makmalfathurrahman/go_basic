package main

import (
	"fmt"
	"go_basic/0031-package_initialization/database"
	_ "go_basic/0031-package_initialization/internal"
)

func main() {
	var database string = database.GetDatabase()

	fmt.Println(database)
}
