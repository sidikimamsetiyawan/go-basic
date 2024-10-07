package main

import (
	"fmt"
)

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	}
	return map[string]string{"name": name}
}

func main() {

	user := newMap("")

	if user == nil {
		fmt.Println("Kosong")
	} else {
		fmt.Println(user)
	}

	// Default Nil
	var s []int
	fmt.Println("Slice", s)
	fmt.Println("Slice == nil =>", s == nil)

	var m map[string]int
	fmt.Println("Map", m)
	fmt.Println("Map == nil =>", m == nil)

	var i interface{}
	fmt.Println("Interface : ", i)

	var ptr *int
	fmt.Println("Pointer : ", ptr)

	var ch chan int
	fmt.Println("Channel : ", ch)

	// Defaul Not Nil
	var _int int
	fmt.Println("Integer : ", _int)

	var str string
	fmt.Println("String : ", str)
}
