package main

import "fmt"

func main() {
	fmt.Println("--- Buffered Channel ---")

	// Create a buffered channel of type string with a capacity of 2
	data := make(chan string, 2)

	// Send two values into the buffered channel
	// These operations do not block because the channel has enough capacity to hold both values.
	data <- "Hello"
	data <- "World"

	// Receive two values from the channel
	// Since both values are already in the buffer, these operations also do not block.
	firstPrint, secondPrint := <-data, <-data

	// Print the received values
	fmt.Println(firstPrint, secondPrint)
}
