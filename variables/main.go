package main

import "fmt"

var x int = 1

// y := 2  // Tidak Bisa DILakukan DiLuar Fungsi

func main() {
	var fullName string = "Sidik Imam Setiyawan"
	var firstName = "Sidik"
	var lastName string

	var a string
	var b int
	var c bool

	// Deklarasi multiple variable
	var d, e, f, g int = 1, 2, 3, 4

	// Deklarasi multiple variable dalam blok
	var (
		address    string = "Jl. Sekeretaris"
		city              = "Tangerang"
		postalCode int    = 15142
	)

	lastName = "Imam Setiyawan"
	myAge := 31

	fmt.Println("fullName : ", fullName)
	fmt.Println("firstName : ", firstName)
	fmt.Println("lastName : ", lastName)
	fmt.Println("myAge : ", myAge)

	fmt.Println("x : ", x)
	// fmt.Println("y : ", y)
	fmt.Println()
	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)

	fmt.Println()
	fmt.Println("d : ", d)
	fmt.Println("e : ", e)
	fmt.Println("f : ", f)
	fmt.Println("g : ", g)

	fmt.Println()
	fmt.Println("address : ", address)
	fmt.Println("city : ", city)
	fmt.Println("postalCode : ", postalCode)

}
