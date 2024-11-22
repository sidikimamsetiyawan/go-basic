package main

import (
	"fmt"
	"sort"
	"strings"
)

// InArray function checks if elements of array1 are substrings of any element in array2
func InArray(array1 []string, array2 []string) []string {
	// Initialize an empty slice to store the result
	r := []string{}
	// Create a map to keep track of unique elements
	unique := make(map[string]bool)
	// Iterate over each element in array1
	for _, x := range array1 {
		// For each element in array1, iterate over each element in array2
		for _, y := range array2 {
			// Check if the element from array1 is a substring of the element from array2
			if strings.Contains(y, x) {
				// If the element is not already in the result slice, add it
				if !unique[x] {
					r = append(r, x)
					// Mark the element as added to ensure uniqueness
					unique[x] = true
				}
			}
		}
	}
	// Sort the result slice in lexicographical order
	sort.Strings(r)
	return r
}

func main() {
	// Define the first array of strings
	a1 := []string{"arp", "live", "strong"}

	a2 := []string{"lively", "alive", "harp", "sharp", "armstrong"}

	// Define the first array of strings
	fmt.Println(InArray(a1, a2))

	a3 := []string{"tarp", "mice", "bull"}

	a4 := []string{"lively", "alive", "harp", "sharp", "armstrong"}

	// Define the first array of strings
	fmt.Println(InArray(a3, a4))
}
