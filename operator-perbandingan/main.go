package main

import "fmt"

func main() {

	var x uint8 = 10
	var y uint8 = 5

	fmt.Println()
	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x > y)
	fmt.Println(x < y)
	fmt.Println(x >= y)
	fmt.Println(x <= y)

	var text1 = "Hello"
	var text2 = "Hello"
	var text3 = "World"

	var result1 bool = text1 == text2
	var result2 bool = text1 == text3

	fmt.Println()
	fmt.Println(result1)
	fmt.Println(result2)

}
