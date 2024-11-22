package main

import "fmt"

func main() {
	// Create a channel of type bool
	c := make(chan bool)

	// Attempt to send a value into the channel
	c <- true

	// This line will never execute because the program will block at the previous line
	// The send operation (c <- true) is blocking, and there is no receiver to consume the value.
	// As a result, the program deadlocks here.
	fmt.Println("This line will never be printed.")
}
