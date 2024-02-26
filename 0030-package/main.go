package main

import (
	"fmt"
	helper "go_basic/0030-package/helper"
)

func main() {
	hello := helper.SayHello("World!")

	fmt.Println(hello)
}
