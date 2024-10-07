package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) GetName() string {
	return u.Name
}

func (u User) GetEmail() string {
	return u.Email
}

func main() {

	var user User

	user.Name = "Sidik"
	user.Email = "sidikimam@gmail.com"
	user.Age = 31

	fmt.Println(user.GetName())
	fmt.Println(user.GetEmail())
}
