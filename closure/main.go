package main

import (
	"fmt"
)

func main() {

	name := "Sidik Imam"

	displayName := func() {
		// Anggap saja tidak sengaja mengubah isi variable
		name = "Shiqiyah"
		fmt.Println(name)
	}

	displayName()
}
