package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("i ke : ", i)
	}

	// Nested loop
	fmt.Println()
	a := [2]string{"jus", "soda"}
	b := [3]string{"jambu", "mangga", "alpukat"}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			fmt.Println(a[i], b[j])
		}
	}

	cars := [3]string{"Toyota", "Honda", "Nissan"}
	for index, car := range cars {
		fmt.Println(index, car)
	}

	var person = map[string]string{
		"name": "Sidik",
		"job":  "backend developer",
	}
	fmt.Println()
	for key, value := range person {
		fmt.Println(key, value)
	}
}
