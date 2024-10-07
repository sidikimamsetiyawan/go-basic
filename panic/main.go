package main

import (
	"fmt"
)

func division(num1, num2 int) {
	if num2 == 0 {
		panic("Tidak dapat membagi angka dengan 0")
	} else {
		result := num1 / num2
		fmt.Println("Result : ", result)
	}
}

func main() {

	division(3, 0)
	// division(3, 1)

}
