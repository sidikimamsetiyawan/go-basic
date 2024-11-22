package main

import "fmt"

func main() {
	// Create a buffered channel of type bool with a capacity of 5
	c := make(chan bool, 5)

	// Send a value into the buffered channel
	// Since the channel has a capacity of 5, this operation does not block.
	c <- true

	// This line will execute because the send operation above is non-blocking.
	fmt.Println("This line be printed.")
}
