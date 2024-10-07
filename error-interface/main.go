package main

import (
	"errors"
	"fmt"
)

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Tidak dapat melakukan pembagian dengan nol")
	} else {
		result := x / y
		return result, nil
	}
}

func main() {
	var exampleError error = errors.New("Examples Error !!")
	fmt.Println(exampleError.Error())

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}
