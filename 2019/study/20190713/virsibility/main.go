package main

import (
	"fmt"
	"virsibility/users"
)

func main() {
	var u users.User
	//var a users.address

	fmt.Printf("%#v\n", u)
	fmt.Println(u.ID)
	fmt.Println(u.Name)
}
