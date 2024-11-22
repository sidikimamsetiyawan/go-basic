package main

import (
	"fmt"
	"strings"
)

func roundUpDivision(numerator, denominator int) int {
	if numerator%denominator == 0 {
		return numerator / denominator
	}
	return (numerator / denominator) + 1
}

func Encode(s string, n int) string {
	// code
	if n <= 1 {
		return s
	}

	// Create a slice to hold each rail's characters
	rail := make([]string, n)
	currentRail := 0
	direction := 1 // 1 for down, -1 for up

	/*
		n the railFenceEncode function, the variable currentRail behaves in a zigzag pattern like 0, 1, 2, 1, 0, 1, 2, ... because of the logic used to traverse the "rails."
	*/

	for _, char := range s {
		// Append character to the current rail
		rail[currentRail] += string(char)
		fmt.Println("cek-", currentRail, " - ", direction, " : ", string(char))
		// Move to the next rail
		currentRail += direction

		// Change direction if we're at the top or bottom rail
		if currentRail == 0 || currentRail == n-1 {
			direction *= -1
			fmt.Println("AAAAAAAAA", rail)
		}
	}
	fmt.Println(rail)
	// Join all rails into a single encoded string
	return strings.Join(rail, "")

}

func Decode(s string, n int) string {
	if n <= 1 {
		return s
	}

	// Calculate the zigzag pattern length
	lengthData := len(s)
	pattern := make([]int, lengthData)
	currentRail := 0
	direction := 1

	for i := 0; i < lengthData; i++ {
		pattern[i] = currentRail
		currentRail += direction
		if currentRail == 0 || currentRail == n-1 {
			direction *= -1
		}
	}

	// Create slices for each rail
	railStrings := make([][]rune, n)
	for i := 0; i < n; i++ {
		railStrings[i] = make([]rune, 0)
	}

	// Distribute characters from s to each rail based on pattern
	cipherIndex := 0
	for railNum := 0; railNum < n; railNum++ {
		for i := 0; i < lengthData; i++ {
			if pattern[i] == railNum {
				railStrings[railNum] = append(railStrings[railNum], rune(s[cipherIndex]))
				cipherIndex++
			}
		}
	}

	// Reconstruct original message
	var decodedMessage strings.Builder
	for i := 0; i < lengthData; i++ {
		decodedMessage.WriteRune(railStrings[pattern[i]][0])
		railStrings[pattern[i]] = railStrings[pattern[i]][1:]
	}

	return decodedMessage.String()
}

func main() {
	// fmt.Println(Encode("WEAREDISCOVEREDFLEEATONCE", 3))
	fmt.Println(Decode("WECRLTEERDSOEEFEAOCAIVDEN", 3))

}
