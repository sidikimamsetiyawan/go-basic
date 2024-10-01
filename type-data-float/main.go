package main

import "fmt"

func main() {

	var a float32 = 125.55
	var b float32 = 3.4e+38 // maksimum value of float32 3.4e+38 atau 3.4 * 10^38

	// var c uint = - 600		// cannot use -600 (untyped int constant) as uint value in variable declaration (overflows)
	var c = 125.55 // Default of float is float64
	var d float32  // if don't added value have default 0 values

	fmt.Println()
	fmt.Printf("Type Data: %T, value: %v\n", a, a)
	fmt.Printf("Type Data: %T, value: %v\n", b, b)

	fmt.Println()
	fmt.Printf("Type Data: %T, value: %v\n", c, c)
	fmt.Printf("Type Data: %T, value: %v\n", d, d)

}
