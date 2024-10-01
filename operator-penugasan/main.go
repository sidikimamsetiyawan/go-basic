package main

import "fmt"

func main() {

	var a uint8 = 3
	var b, c, d, e, f = 5, 10, 15, 20, 25

	b += 5 // b = b + 5
	c -= 3 // c = c - 5
	d *= 5 // d = d * 5
	e /= 5 // e = e / 5
	f %= 5 // f = f % 5

	fmt.Println()
	fmt.Println("Nilai a : ", a)
	fmt.Println("Nilai b : ", b)
	fmt.Println("Nilai c : ", c)
	fmt.Println("Nilai d : ", d)
	fmt.Println("Nilai e : ", e)
	fmt.Println("Nilai f : ", f)

}
