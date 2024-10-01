package main

import "fmt"

const LENGTH int = 11
const WIDTH = 5

// multiple constants
const (
	X int = 10
	Y     = 3.14
	Z     = "Hello World"
)

// y := 2  // Tidak Bisa DILakukan DiLuar Fungsi

func main() {

	const PI = 3.14
	// PI = 10 //cannot assign to PI (neither addressable nor a map index expression)

	fmt.Println("LENGTH : ", LENGTH)
	fmt.Println("WIDTH : ", WIDTH)
	// fmt.Println("PI : ", PI)
	fmt.Println()
	fmt.Println("X : ", X)
	fmt.Println("Y : ", Y)
	fmt.Println("Z : ", Z)

}
