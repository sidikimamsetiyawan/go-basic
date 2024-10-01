package main

import "fmt"

func main() {

	var a1 = [3]int{100, 200, 399}
	a2 := [4]int{4, 5, 6, 7}

	var a3 = [...]int{11, 22, 33}
	a4 := [...]int{11, 12, 13, 14}

	var a5 = [2]string{"Go", "C#"}

	prices := [3]int{10000, 20000, 30000}

	a6 := [4]int{}           // Tidak diinisialisasi, maka akan terisi default 0 [0 0 0 0]
	a7 := [4]int{0, 1}       // Diinisialisasi sebagian
	a8 := [4]int{0, 1, 2, 3} // Diinisialisasi semua

	fmt.Println()
	fmt.Println(a1)
	fmt.Println(a2)

	fmt.Println()
	fmt.Println(a3)
	fmt.Println(a4)

	fmt.Println()
	fmt.Println(a5)

	fmt.Println()
	fmt.Println(prices)
	fmt.Println(prices[0])
	fmt.Println(prices[2])

	prices[2] = 25000
	fmt.Println()
	fmt.Println(prices)

	fmt.Println()
	fmt.Println(a6)
	fmt.Println(a7)
	fmt.Println(a8)

	fmt.Println()
	fmt.Println(len(prices))

}
