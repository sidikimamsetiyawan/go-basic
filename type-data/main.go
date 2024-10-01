package main

import "fmt"

func main() {

	var isMarried bool = true               // Boolean
	var myAge int = 31                      // Integer
	var pi float32 = 3.14                   // FLoationg Point
	var myFullName = "Sidik Imam Setiyawan" // String

	// without value have a return default value

	var defaultIsMarried bool    // Boolean - false
	var defaultMyAge int         // Integer - 0
	var defaultPi float32        // FLoationg Point - 0
	var defaultMyFullName string // String - Empty String

	fmt.Println()
	fmt.Println("isMarried : ", isMarried)
	fmt.Println("myAge : ", myAge)
	fmt.Println("pi : ", pi)
	fmt.Println("myFullName : ", myFullName)

	fmt.Println()
	fmt.Println("defaultIsMarried : ", defaultIsMarried)
	fmt.Println("defaultMyAge : ", defaultMyAge)
	fmt.Println("defaultPi : ", defaultPi)
	fmt.Println("defaultMyFullName : ", defaultMyFullName)

}
