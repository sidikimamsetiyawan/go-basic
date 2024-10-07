package main

import (
	"fmt"
)

type User struct {
	Email string
}

func (u *User) RewriteInfo() {
	u.Email = "Alamat Email : " + u.Email
}

func (u *User) RewriteInfoAndReturn() string {
	u.Email = "Ubah Email : " + u.Email
	return u.Email
}

func main() {
	var user User
	user.Email = "user@gmail.com"

	user.RewriteInfo()      // call the function method, for modify the value of user.Email
	fmt.Println(user.Email) // Print struct user

	getEmail := user.RewriteInfoAndReturn()
	fmt.Println()
	fmt.Println(getEmail)

}
