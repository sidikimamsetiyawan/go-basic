package main

import "fmt"

// Deklarasi Interfase
type Greeting interface {
	Greet() string
}

// Deklarasi struct
type Person struct {
	Name string
}

// Implementasi - Argument struct (Person)
// Create function name : Greet
func (p Person) Greet() string {
	return p.Name
}

// function sayHello dengan argument interface
// argument harus sesuai dengan function implementasi
func SayHello(g Greeting) {
	fmt.Println("Hello my name is ", g.Greet())
}

func main() {

	person := Person{Name: "Sidik"}

	SayHello(person)
}
