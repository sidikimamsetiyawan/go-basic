package main

import "fmt"

func countRecursive(x int) int {
	if x == 11 {
		return 0 // Part or stop recursive
	}
	fmt.Println(x)
	return countRecursive(x + 1) // Part or recursive
}

func factorialLoop(y int) int {
	result := 1
	for i := y; i > 0; i-- {
		fmt.Println(result, " * ", i)
		result *= i
	}
	return result
}
func factorialRecursive(z int) int {
	if z == 0 {
		return 1
	}
	return z * factorialRecursive(z-1)
}

func main() {
	countRecursive(1)

	// Factorial with loop function
	loop := factorialLoop(5)
	fmt.Println(loop)

	// Example case : Factorial
	// 5! = 5 * 4 * 3 * 2 * 1 = 120
	x := factorialRecursive(5)
	fmt.Println("Example case : Factorial")
	fmt.Println(x)

}
