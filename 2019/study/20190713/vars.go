package main

import "fmt"

type Address struct {
	Region string
	Street string
	No     string
}

type User struct {
	ID   int
	Name string
	Addr Address
}

func changePoint(u *User) {
	u.Name = "AAAA"
}

func main() {
	me := User{}
	me2 := me
	me2.Name = "danran"

	fmt.Printf("%#v\n", me)
	fmt.Printf("%#v\n", me2)

	changePoint(&me2)
	fmt.Printf("%#v\n", me2)

}
