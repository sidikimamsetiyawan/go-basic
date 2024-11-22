package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Function to check if a character is a vowel
func isVowel(ch rune) bool {
	vowels := "aeiouAEIOU"
	return strings.ContainsRune(vowels, ch)
}

// Function to check if a character is a consonant
func isConsonant(ch rune) bool {
	// Check if it’s a letter and not a vowel
	return (ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z') && !isVowel(ch)
}

func EncryptThis(text string) string {
	if len(text) == 0 {
		return "String is empty."
	}

	words := strings.Split(text, " ")
	var encryptedWords []string

	for _, word := range words {
		if len(word) == 0 {
			continue
		}

		// Get ASCII code of the first character
		asciiCode := strconv.Itoa(int(word[0]))

		// Build the encrypted word with ASCII code as the start
		encryptedWord := asciiCode
		// completeWords := ""
		// oddWords := ""

		// If the word has more than two characters, swap the second and last
		runes := []rune(word)
		fmt.Println("runes : ", runes)
		fmt.Println("len(runes) : ", len(runes))
		if len(runes) > 2 {
			encryptedWord += string(runes[len(runes)-1]) + string(runes[2:len(runes)-1]) + string(runes[1])
			// fmt.Println("Cek isConsonant-", string(runes[1]), " : ", isConsonant(runes[1]))
			// for i := len(runes) - 1; i != 0; i-- {
			// 	fmt.Println("isConsonant-", string(runes[i]), " : ", isConsonant(runes[i]))
			// 	// if len(runes)-1%2 == 0 {
			// 	if isConsonant(runes[1]) {
			// 		if i%2 == 0 {
			// 			completeWords += string(runes[i])
			// 		} else {
			// 			oddWords += string(runes[i])
			// 		}
			// 		if i == 1 {
			// 			encryptedWord += completeWords
			// 			encryptedWord += oddWords
			// 		}
			// 	} else {
			// 		encryptedWord += string(runes[i])
			// 	}
			// }
			// for i := len(runes) - 1; i != 0; i-- {
			// 	if i%2 == 0 {
			// 		completeWords += string(runes[i])
			// 	} else {
			// 		oddWords += string(runes[i])
			// 	}

			// 	if i == 1 {
			// 		encryptedWord += completeWords
			// 		encryptedWord += oddWords
			// 	}
			// 	// poke => eokp
			// 	// 1 2 3 4 => 4 2 3 1
			// 	// ise => esi
			//  // 1 2 3 => 3 2 1
			// // ived => dvei
			// // 1 2 3 4 => 4 2 3 1
			// 	fmt.Println("runes-", i, " : ", string(runes[i]))
			// 	// encryptedWord += string(runes[i])
			// }
		} else if len(runes) == 2 {
			// If only two characters, add the last character directly
			encryptedWord += string(runes[1])
		}
		fmt.Println("encryptedWord : ", encryptedWord)
		encryptedWords = append(encryptedWords, encryptedWord)
		fmt.Println("encryptedWords : ", encryptedWords)
	}

	return strings.Join(encryptedWords, " ")
}

func main() {
	// {
	// 	{"hello world", "104olle 119drlo"}, // Normal case
	// 	{"", "String is empty."},           // Empty string case
	// 	{"a", "97"},                        // Single character case
	// 	{"ab", "97b"},                      // Two characters case
	// 	{"abc", "97cb"},                   // Three characters case
	// 	{"abcd", "97dcb"},                  // Four characters case
	// 	{"abcde", "97edcba"},               // Five characters case
	// 	{"a b c", "97 98 99"},              // Multiple single characters
	// 	{"hello  ", "104olle"},         // Trailing spaces
	// 	{"  hello", "104olle"},             // Leading spaces
	// 	{"  ", ""},                         // Only spaces
	// 	{"a b c d e", "97 98 99 100 101"},  // Multiple single characters
	// 	{"test case", "116tase 99esa"},       // Normal case with two words
	// }
	// str := "hello world"
	str := "spoke"
	// 65 119esi 111dl 111lw 108devi 105n 97n 111ka
	// 65 119esi 111dl 111lw 108dvei 105n 97n 111ka
	// 84eh 109ero 104e 115wa 116eh 108sse 104e 115ekop
	// 84eh 109ero 104e 115wa 116eh 108sse 104e 115eokp - spoke : 115 - e - o - k - p
	fmt.Println(EncryptThis(str)) // Expected output: "104olle 119drlo"
}
