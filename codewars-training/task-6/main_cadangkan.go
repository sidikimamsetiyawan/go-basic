package main_cadangkan

import (
	"fmt"
	"math/big"
)

// intPow calculates the power of a base raised to an exponent using an efficient method.
func intPow(base, exponent int) int {
	result := 1 // Start with result = 1, as any number raised to power 0 is 1.

	// Loop until the exponent becomes 0.
	for exponent > 0 {
		// If the exponent is odd, multiply the current result by the base.
		if exponent%2 == 1 { // Check if the exponent is odd
			result *= base
		}

		// Square the base for the next power level.
		base *= base

		// Halve the exponent to reduce the number of multiplications.
		exponent /= 2
	}

	// Return the computed result.
	return result
}

// LastDigit computes the last digit of a large exponentiation sequence.
func LastDigit(as []int) int {
	res := big.NewInt(0) // Initialize a big integer to hold the result.
	pow := 0             // Variable to store the current exponent.

	// If the input slice is empty, return 1 (as 0^0 is conventionally defined as 1).
	if len(as) == 0 {
		return 1
	}

	// Iterate through the slice from the last element to the first.
	for i := len(as) - 1; i >= 0; i-- {
		if len(as)-1 == i {
			// For the last element, set it as the initial exponent.
			pow = as[i]
		} else {
			// Calculate the current power using the previous exponent.
			res = new(big.Int).Exp(big.NewInt(int64(as[i])), big.NewInt(int64(pow)), nil)
			// Update the exponent for the next iteration using intPow.
			pow = intPow(as[i], pow)
		}
	}

	// Convert the result to a string to extract the last digit.
	resStrings := res.String()
	// Get the last character of the result string, convert it to an integer.
	result := int(resStrings[len(resStrings)-1] - '0')
	return result // Return the last digit of the result.
}

func main_cadangkan() {
	// Sample data for testing the LastDigit function.
	sampleData := []int{0, 0}
	// Print the last digit of the exponentiation result.
	fmt.Println(LastDigit(sampleData)) // Expected output: 1 (since 3^4^5 ends with 1)
}
