package main

import "fmt"

func main() {

	fmt.Println(10 + 4)
	fmt.Println(10 - 4)
	fmt.Println(10 * 4)
	fmt.Println(10 / 2)
	fmt.Println(10 % 4)

	var a, b, c = 10, 4, 2

	fmt.Println()
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / c)
	fmt.Println(a % b)

	var d = a + b + c
	fmt.Println()
	fmt.Println(d)

	i := 5
	i++

	j := 5
	j--

	fmt.Println()
	fmt.Println(i)
	fmt.Println(j)

}
