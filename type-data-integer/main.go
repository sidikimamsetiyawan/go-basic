package main

import "fmt"

func main() {

	var a int = 600
	var b int = -4600

	// var c uint = - 600		// cannot use -600 (untyped int constant) as uint value in variable declaration (overflows)
	var c uint = 600
	var d uint = 4600

	fmt.Println()
	fmt.Printf("Type Data: %T, value: %v\n", a, a)
	fmt.Printf("Type Data: %T, value: %v\n", b, b)

	fmt.Println()
	fmt.Printf("Type Data: %T, value: %v\n", c, c)
	fmt.Printf("Type Data: %T, value: %v\n", d, d)

}
