package main

import (
	"fmt"
)

func main() {
	type itsString string
	type itsInt int
	type itsBoolean bool

	var text itsString = "Hello World"
	var number itsInt = 10
	var _boolean itsBoolean = true

	fmt.Println()
	fmt.Println(text)
	fmt.Println(number)
	fmt.Println(_boolean)
}
