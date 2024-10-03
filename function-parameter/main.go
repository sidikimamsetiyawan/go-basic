package main

import "fmt"

func getName(firstName string) {
	fmt.Println("Get Data Firtsname : ", firstName)
}

func getData(fullName string, age int) {
	fmt.Println("My name is ", fullName, " . I'm ", age, "years old")
}
func main() {
	getName("Sidik")
	getData("Sidik Imam Setiyawan", 31)
}
