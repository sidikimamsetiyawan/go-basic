package main

import "fmt"

func printType(i interface{}) {
	fmt.Printf("%v, %T\n", i, i)
}

func main() {

	printType(10)
	fmt.Println()
	printType("Hello World")
	fmt.Println()
	printType(true)
}
