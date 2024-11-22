package main

import "fmt"

func main() {
	// Create an unbuffered channel of type bool
	c := make(chan bool)

	// Start a goroutine to receive from the channel
	go func() {
		<-c // This goroutine waits to receive a value from the channel 'c'.
	}()

	// Send a value into the channel
	// This will not block because the goroutine above is ready to receive the value.
	c <- true

	// Since the send operation is completed, the program continues and prints the line below.
	fmt.Println("This line is printed.")
}
