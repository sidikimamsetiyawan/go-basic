package main

import "fmt"

func main() {
	// Create an unbuffered channel of type string
	data := make(chan string)

	// Call the function printRoutine and pass the channel as an argument
	printRoutine(data)

	// Attempt to receive a value from the channel
	word := <-data

	// Print the value received from the channel
	fmt.Println(word)
}

func printRoutine(data chan string) {
	// Send a value into the channel
	// This line will block because the main function has not yet started receiving from the channel.
	data <- "From other routine."
}
