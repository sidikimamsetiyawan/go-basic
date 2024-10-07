package main

import (
	"fmt"
)

func txtNumber() {
	fmt.Println("Tiga")
}

func main() {

	defer txtNumber()

	fmt.Println("Satu")
	fmt.Println("Dua")

}
