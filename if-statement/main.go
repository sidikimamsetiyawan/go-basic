package main

import "fmt"

func main() {
	number := 5

	if number == 5 {
		fmt.Println()
		fmt.Println("Number is 5")
	}

	if number%2 == 0 {
		fmt.Println()
		fmt.Println("Genap")
	} else {
		fmt.Println()
		fmt.Println("Ganjil")
	}

	nilai := 85
	absen := 90

	if nilai > 90 && absen > 95 {
		fmt.Println()
		fmt.Println("Sangat Baik")
	} else if nilai > 90 && absen > 85 {
		fmt.Println()
		fmt.Println("Baik")
	} else {
		fmt.Println()
		fmt.Println("Cukup")
	}

	// Nested IF
	num := 100

	if num >= 80 {
		fmt.Println()
		fmt.Println("Num lebih dari 80")
		if num > 90 {
			fmt.Println("Num juga lebih dari 90")
		}
	} else {
		fmt.Println()
		fmt.Println("Num kurang dari 80")
	}

}
