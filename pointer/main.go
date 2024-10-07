package main

import (
	"fmt"
)

func main() {
	// Part Pass By Value
	var a string = "Sidik"
	var b string = a

	b = "Imam"

	fmt.Println(a)
	fmt.Println(b)

	// Part Pointer - Pass By Reference
	// ( * ) bisa dijadikan pointer atau operator yang bisa mengambil value
	var i string = "Sidik"
	var j *string = &i

	*j = "Imam"

	fmt.Println()
	fmt.Println(i)
	fmt.Println(j)  // Alamat
	fmt.Println(*j) // Value
}
