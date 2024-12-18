package main

import "fmt"

type FullUser struct {
	UserName  string
	UserEmail string
}

type SimpleUser struct {
	UserName string
}

func RemoveIndex(s []int, index int) []int {
	// s[]int : arguments of data slice
	// index : arguments for index position to want to be deleted

	ret := make([]int, 0)
	// New variable ret - For the new data slice without index deleted
	ret = append(ret, s[:index]...)
	// append data to variable ret, from 0 index to arguments index ( Not include index deleted)
	return append(ret, s[index+1:]...)
	// append data to variable ret, from argument index + 1
}

func UpdateIndex(s []int, index int, value int) []int {
	retvar := make([]int, 0)
	length_data := len(s)
	retvar = append(retvar, s[:length_data]...)
	retvar[index] = value
	return retvar
}

func filter(fu []FullUser, su []SimpleUser) (out []FullUser) {
	f := make(map[string]struct{}, len(su))
	for _, u := range su {
		f[u.UserName] = struct{}{}
	}
	for _, u := range fu {
		if _, ok := f[u.UserName]; ok {
			out = append(out, u)
		}
	}
	return
}

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

	// Remove Slice Data
	all := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("all1: ", all) //[0 1 2 3 4 5 6 7 8 9]
	removeIndex := RemoveIndex(all, 5)

	fmt.Println("all2: ", all)                //[0 1 2 3 4 5 6 7 8 9]
	fmt.Println("removeIndex: ", removeIndex) //[0 1 2 3 4 6 7 8 9]

	// Update Index [0]
	updateIndex := UpdateIndex(all, 0, 999)
	fmt.Println("all3: ", all)                //[0 1 2 3 4 5 6 7 8 9]
	fmt.Println("updateIndex: ", updateIndex) //[999 1 2 3 4 6 7 8 9]

	fmt.Println()
	fmt.Println()
	fmt.Println("Hello, playground")

	manyFullUsers := []FullUser{{"foo", "foo@jawohl.com"},
		{"bar", "bar@jawohl.com"},
		{"baz", "baz@jawohl.com"}}

	manySimpleUsers := []SimpleUser{{"foo"}, {"bar"}}

	fmt.Println(manyFullUsers)
	fmt.Println(manySimpleUsers)
	fmt.Println(filter(manyFullUsers, manySimpleUsers))

}
