package main

import "fmt"

func main() {
	//Array
	itsArray1 := [5]int{1, 2, 3, 4, 5}
	itsArray2 := [...]int{1, 2, 3, 4, 5}

	//Slice
	var itsSlice1 = []int{1, 2, 3, 4, 5, 6}
	itsSlice2 := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jum'at"}

	fmt.Println("Part Array")
	fmt.Println(itsArray1)
	fmt.Println(itsArray2)
	fmt.Println("Part Slice")
	fmt.Println(itsSlice1)
	fmt.Println(itsSlice2)
	fmt.Println(len(itsSlice1))
	fmt.Println(cap(itsSlice2))

	fmt.Println("Create Slice From Array")
	newArray := [8]int{11, 12, 13, 14, 15, 16, 17, 18}
	itsSlice3 := newArray[2:4]

	fmt.Println(itsSlice3)
	fmt.Println(len(itsSlice3))
	fmt.Println(cap(itsSlice3))

	fmt.Printf("Ini Slice 3 :%d\n ", cap(itsSlice3)) // 13, 14, 15, 16, 17, 18

	itsSlice4 := newArray[2:]

	fmt.Println(itsSlice4)
	fmt.Println(len(itsSlice4))
	fmt.Println(cap(itsSlice4))

	itsSlice5 := newArray[:8]

	fmt.Println(itsSlice5)
	fmt.Println(len(itsSlice5))
	fmt.Println(cap(itsSlice5))

	itsSlice6 := newArray[:]

	fmt.Println(itsSlice6)
	fmt.Println(len(itsSlice6))
	fmt.Println(cap(itsSlice6))

	newSlice1 := make([]int, 6, 6)
	// length panjang karakter
	// kapasitas jumlah kapasitas untuk menambahkan bisa pakai append
	newSlice1[0] = 100
	newSlice1[1] = 200
	newSlice1[2] = 300
	newSlice1[3] = 400
	newSlice1[4] = 500
	newSlice1[5] = 600

	fmt.Println(newSlice1)
	fmt.Println(len(newSlice1))
	fmt.Println(cap(newSlice1))

}
