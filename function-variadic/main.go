package main

import (
	"fmt"
)

func sum(nums ...int) int {
	// parameter harus diletakan di akhir
	total := 0

	for _, num := range nums {
		total += num
	}
	return total
}

func summary(param1 string, values ...int) (studentName string, value int) {

	value = 0

	for _, num := range values {
		value += num
	}
	return param1, value
}

func main() {
	sumData := sum(1, 2, 3, 4, 5)
	fmt.Println(sumData)

	studentName, value := summary("Sidik Imam Setiyawan", 1, 2, 3, 4, 5)

	fmt.Println(studentName)
	fmt.Println(value)
	// Variadic Function Slice Parameter
	nums := []int{10, 20, 30, 40, 50}
	total := sum(nums...)
	fmt.Println(total)
}
