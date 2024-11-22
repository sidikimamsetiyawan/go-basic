package main

import (
	"fmt"
	"math"
	"strconv"
)

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

// LastDigit computes the last digit of a large exponentiation sequence using the math/big package.
func LastDigit(as []int) int {
	// Initialize a big integer to hold the result.
	res := float64(0)
	res8 := 0
	pow := float64(1)

	// If the input slice is empty, return 1 (since conventionally 0^0 is 1).
	if len(as) == 0 {
		return 1
	} else if len(as) == 2 {
		// res = new(big.Int).Exp(float64(int64(as[0])), float64(int64(as[1])), nil)
		res8 = intPow(as[0], as[1])
		res = float64(int64(res8))
	} else {

		// Iterate through the slice from the last element to the first.
		for i := len(as) - 1; i >= 0; i-- {
			base := float64(int64(as[i]))

			if i == len(as)-1 {
				// For the last element, set it as the initial exponent.
				pow = base
			} else {
				// Calculate the current power using the previous exponent.

				res = math.Pow(base, pow)

				pow = res
				fmt.Println("res", res)
			}
		}
	}
	fmt.Println("result", res)
	// Convert the result to a string to extract the last digit.
	str := strconv.Itoa(int(math.Round(res)))
	// resStrings := resddd.String()
	lastDigit := int(str[len(str)-1] - '0')

	return lastDigit
}

func main() {
	// Sample data for testing the LastDigit function.
	sampleData := []int{7, 6, 21}
	// Print the last digit of the exponentiation result.
	fmt.Println(LastDigit(sampleData)) // Expected output: 1 (since 3^4^5 ends with 1)
}
