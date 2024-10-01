package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3}
	// Access the element slice
	fmt.Println(numbers)
	fmt.Println(numbers[0])
	fmt.Println(numbers[2])

	// Update the element slice
	numbers[2] = 10
	fmt.Println()
	fmt.Println(numbers[2])

	// Add the element slice
	numbers = append(numbers, 11, 12)
	fmt.Println()
	fmt.Println(numbers)

	// Copy slice
	sliceData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	newSliceData := make([]int, len(sliceData), cap(sliceData))
	copy(newSliceData, sliceData)
	fmt.Println()
	fmt.Println(sliceData)
	fmt.Println(newSliceData)

}
