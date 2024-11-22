package main

import (
	"fmt"
)

func sendData(ch chan int) {
	ch <- 42 // Kirim data ke channel
}

func main() {
	ch := make(chan int) // Membuat channel
	go sendData(ch)      // Menjalankan goroutine
	value := <-ch        // Menerima data dari channel
	fmt.Println(value)   // Output: 42
}
