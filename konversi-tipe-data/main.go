package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int8 = 10
	var b int16 = int16(a)
	var newint = "100"
	var myint, _ = strconv.Atoi(newint)

	fmt.Println()
	fmt.Printf("Tipe data a: %T, nilai a: %v\n", a, a)
	fmt.Printf("Tipe data b: %T, nilai b: %v\n", b, b)
	fmt.Printf("Tipe data myint: %T, nilai myint: %v\n", myint, myint)

	var c uint32 = 255
	var d uint64 = uint64(c)

	var e float32 = 20.5
	var f float64 = float64(e)

	fmt.Println()
	fmt.Printf("Tipe data c: %T, nilai c: %v\n", c, c)
	fmt.Printf("Tipe data d: %T, nilai d: %v\n", d, d)

	fmt.Println()
	fmt.Printf("Tipe data e: %T, nilai e: %v\n", e, e)
	fmt.Printf("Tipe data f: %T, nilai f: %v\n", f, f)

	// Byte to string
	var name string = "Sidik"
	var charA byte = name[0]               // Huruf S dalam byte di konversikan 83
	var charAString string = string(charA) // Byte to string

	fmt.Println()
	fmt.Printf("Tipe data name: %T, nilai name: %v\n", name, name)
	fmt.Printf("Tipe data charA: %T, nilai charA: %v\n", charA, charA)
	fmt.Printf("Tipe data charAString: %T, nilai charAString: %v\n", charAString, charAString)
}
