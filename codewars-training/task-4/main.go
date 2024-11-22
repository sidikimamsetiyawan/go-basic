package main

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {

	arr := strings.Split(in, " ")

	min, _ := strconv.Atoi(arr[0])

	max := min

	for _, number := range arr {
		number_, _ := strconv.Atoi(number)

		if number_ < min {
			min = number_
		}
		if number_ > max {
			max = number_
		}

	}

	result := strconv.Itoa(max) + " " + strconv.Itoa(min)
	return result

}

func main() {
	fmt.Println(HighAndLow("10 20 30 40 50"))
}
