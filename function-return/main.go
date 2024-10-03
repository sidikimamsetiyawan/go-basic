package main

import "fmt"

func getSum(a int, b int) int {
	return a + b
}

func getSum2(x int, y int) (result int) {
	result = x + y
	return
}

func getSum3(i int, j string) (sum int, fullname string) {
	sum = i + i
	fullname = j + " Imam "
	return
}

func getSum4(i int, j string) (sum int, fullname string) {
	sum = i + i
	fullname = j + " Imam "
	return
}

func main() {
	datasummary := getSum(10, 20)
	datasummary2 := getSum2(11, 22)
	x1, x2 := getSum3(10, "Sidik")
	y1, _ := getSum3(10, "Sidik")

	fmt.Println("datasummary : ", datasummary)
	fmt.Println("datasummary2 : ", datasummary2)
	fmt.Println("x1 : ", x1)
	fmt.Println("x2 : ", x2)
	fmt.Println("y1 : ", y1)
}
