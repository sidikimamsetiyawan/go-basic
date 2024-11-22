package main

import (
	"fmt"
	"strconv"
)

// SumMix function takes a slice of mixed types (int and string) and returns the sum of all elements as int64
func SumMix(arr []any) int {
	// Initialize total to store the sum of elements
	var total int64 = 0

	// Iterate over each element in the input slice
	for _, number := range arr {
		// Use type switch to handle different types in the slice
		switch v := number.(type) {
		case int:
			// Handle int type: convert to int64 and add to total
			total += int64(v)
		case string:
			// Handle string type: convert to int64 and add to total
			i, _ := strconv.ParseInt(v, 10, 64)
			total += i
		default:
			// Handle unsupported types
			fmt.Println("Unsupported type")
		}
	}

	// Return the total sum of elements
	return int(total)
}

func main() {
	// Test case: a slice containing both int and string types
	data := []any{9, 3, "7", "3"}

	// Print the result of the SumMix function
	fmt.Println(SumMix(data)) // Output: 15
}
