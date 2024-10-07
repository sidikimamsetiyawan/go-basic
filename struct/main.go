package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {

	var user User

	user.Name = "Sidik"
	user.Email = "sidikimam@gmail.com"
	user.Age = 31

	fmt.Println(user)
	fmt.Println(user.Name)

	var myuser = User{
		Name:  "Imam",
		Email: "imamsidik@gmail.com",
		Age:   30,
	}

	fmt.Println(myuser)
	fmt.Println(myuser.Name)

	var exuser = User{
		"exemployee",
		"employee@gmail.com",
		25,
	}

	fmt.Println(exuser)
	fmt.Println(exuser.Name)
}
