package main

import "fmt"

func main() {

	var firstName, lastName string = "Sidik", "Imam Setiyawan"

	var x string = "Hello"
	var y int = 25

	var numberData = 15.5
	var textData = "My World"

	fmt.Print("firstName : ", firstName)
	fmt.Print("lastName : ", lastName)
	fmt.Print("\n")
	fmt.Print(firstName, "\n")
	fmt.Print(lastName, "\n")
	fmt.Print("\n")
	fmt.Print(firstName, " ", lastName)
	fmt.Print("\n")
	fmt.Println()
	fmt.Println("firstName : ", firstName)
	fmt.Println("lastName : ", lastName)
	fmt.Println()
	fmt.Printf("x value is : %v and type data x is : %T\n", x, x)
	fmt.Printf("y value is : %v and type data y is : %T\n", y, y)
	fmt.Println()
	fmt.Printf("numberData : %v%%\n ", numberData)
	fmt.Printf("textData : %#v\n", textData)

}
