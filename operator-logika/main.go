package main

import "fmt"

func main() {

	var a uint8 = 15

	fmt.Println()
	fmt.Println(a > 10 && a < 5)
	fmt.Println(a > 10 || a < 5)
	fmt.Println(!(a > 10 || a < 5))

}
