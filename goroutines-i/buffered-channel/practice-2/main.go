package main

import "fmt"

func main() {
	c := make(chan int, 2) // Create a buffered channel with a capacity of 2

	c <- 1 // First value is sent
	c <- 2 // Second value is sent

	fmt.Println("Channel is full now.")

	// This line will block because the channel is full
	c <- 3

	// The program will not reach this point until space is freed in the channel
	fmt.Println("This line will not be reached unless a receiver retrieves a value.")
}
