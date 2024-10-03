package main

import (
	"fmt"
)

func main() {
	// Sample Anonymous function
	var getHello = func() {
		fmt.Println("Hello anonymous function!")
	}

	getHello()
	// Sample Anonymous function with parameter
	var getData = func(fullName string, age int) {
		fmt.Println("Nama saya ", fullName, "dan umur saya ", age)
	}

	getData("Sidik Imam Setiyawan", 31)

	// Sample Anonymous function with parameter dan return value
	var getContactNumber = func(x string) string {
		return x
	}

	contactNumber := getContactNumber("08784123xxxx")
	fmt.Println(contactNumber)
}
