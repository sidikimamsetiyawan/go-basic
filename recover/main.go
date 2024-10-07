package main

import (
	"fmt"
)

func division(num1, num2 int) {
	defer handlePanic()
	if num2 == 0 {
		panic("Tidak dapat membagi angka dengan 0")
	} else {
		result := num1 / num2
		fmt.Println("Result : ", result)
	}
}

func handlePanic() {
	err := recover()

	if err != nil {
		fmt.Println("RECOVER", err)
	}
}

func main() {

	division(3, 0) // Test Error
	// division(3, 1) // Test Success

}
