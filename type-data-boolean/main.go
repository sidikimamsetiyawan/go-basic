package main

import "fmt"

func main() {

	var a bool = true // Declare variable with initiate type data
	var b = true      // Declare variable without initiate type data
	var c bool        // Declare variable without value
	d := true         // Declare variable with :=

	fmt.Println()
	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)
	fmt.Println("d : ", d)
	fmt.Printf("%v, %T\n", b, b)

}
