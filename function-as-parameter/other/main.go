package main

import (
	"fmt"
	"strings"
)

// Example Function Typer Declaration
type Formatter func(string) string

// Part Function As Parameter
// function with name : formatter - call function upperCase
func formatCase(txt string, formatter Formatter) {
	// fmt.Println("Cek-1")
	formatted := formatter(txt)
	// fmt.Println("Cek-1-1", formatted)
	fmt.Println(formatted)
}

func upperCase(txt string) string {
	// fmt.Println("Cek-2")
	return strings.ToUpper(txt)
}

func lowerCase(txt string) string {
	return strings.ToLower(txt)
}

func main() {

	formatCase("hello world", upperCase)
	formatCase("LOWER CASE !!! ", lowerCase)

	toUpper := upperCase
	formatCase("New World!!", toUpper)
}
