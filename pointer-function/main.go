package main

import (
	"fmt"
)

type Person struct {
	Name string
	Job  string
}

func rewriteInfo(person *Person) {
	// person.Job = "Programmer"
	person.Name = "Imam"
}

func rewriteInfoAndReturn(person *Person) string {
	// person.Job = "Programmer"
	person.Name = "Imam"
	return person.Name
}

func main() {
	var person = Person{
		Name: "Sidik",
		Job:  "Backend Developer",
	}
	fmt.Println()
	fmt.Println(person)

	rewriteInfo(&person) // Jangan lupa & untuk passing with value

	fmt.Println()
	fmt.Println(person)

	getName := rewriteInfoAndReturn(&person) // Jangan lupa & untuk passing with value

	fmt.Println()
	fmt.Println(getName)
}
