package main

import "fmt"

func main() {
	operator(5, 2)
	fmt.Println("----------")
	augmented(10, 10, 10, 10.0, 10)
	fmt.Println("----------")
	unary(5, 5)
	fmt.Println("----------")
	comparison(5, 2)
	fmt.Println("----------")

}

func operator(a int, b int) {
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)

	var aFloat = float64(a)
	var bFloat = float64(b)
	fmt.Println(aFloat / bFloat)

	fmt.Println(a % b)
}

func augmented(c1 int, c2 int, c3 int, c4 float64, c5 int) {
	c1 += 5
	fmt.Println(c1)

	c2 -= 50
	fmt.Println(c2)

	c3 *= 2
	fmt.Println(c3)

	c4 /= 4.0
	fmt.Println(c4)

	c5 %= 2
	fmt.Println(c5)
}

func unary(d1 int, d2 int) {
	d1++
	fmt.Println(d1)

	d2--
	fmt.Println(d2)
}

func comparison(e1 int, e2 int) {
	fmt.Println(e1 < e2)
	fmt.Println(e1 > e2)
	fmt.Println(e1 <= e2)
	fmt.Println(e1 >= e2)
	fmt.Println(e1 == e2)
	fmt.Println(e1 != e2)
}
