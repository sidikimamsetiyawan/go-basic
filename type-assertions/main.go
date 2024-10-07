package main

import (
	"fmt"
)

func main() {
	var newData interface{} = "Hello World"
	fmt.Println(newData)

	stringData := newData.(string)
	fmt.Println(stringData)

	// Handle type data assertions
	fmt.Println()
	var myData interface{} = "love"
	num, ok := myData.(int)

	if !ok {
		recoverMyData := recover()
		fmt.Println("Type data has been revert ( not integer )", recoverMyData)
	} else {
		fmt.Println("Type data is integer", num)
	}

	// Handle type data assertions with switch
	fmt.Println()

	switch myData.(type) {
	case int:
		fmt.Println("My Data is integer", num)
	case string:
		fmt.Println("My Data is string", num)
	default:
		fmt.Println("My Data is other type data", num)
	}

}
