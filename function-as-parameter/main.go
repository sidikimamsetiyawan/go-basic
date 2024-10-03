package main

import (
	"fmt"
	"strings"
)

// Part Function As Parameter
// function with name : formatter - call function upperCase
func formatCase(txt string, formatter func(string) string) {
	fmt.Println("Cek-1")
	formatted := formatter(txt)
	fmt.Println("Cek-1-1", formatted)
	fmt.Println(formatted)
}

func upperCase(txt string) string {
	fmt.Println("Cek-2")
	return strings.ToUpper(txt)
}

func main() {

	formatCase("hello world", upperCase)

	toUpper := upperCase
	formatCase("New World!!", toUpper)
}
