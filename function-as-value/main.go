package main

import (
	"fmt"
)

func sum(a, b int) int {

	return a + b
}

func main() {

	total := sum
	fmt.Println(total(1, 2))

	result := total(1, 2)
	fmt.Println(result)
}
