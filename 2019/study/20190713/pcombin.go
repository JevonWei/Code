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
	Addr *Address
}

func main() {
	var me User
	fmt.Printf("%#v\n", me)

	me02 := User{
		ID:   1,
		Name: "Jevon",
		Addr: &Address{"shanghai", "shuangqiao", "1139"},
	}
	fmt.Printf("%#v\n", me02)

	me02.Addr.Region = "tang"
	fmt.Println(me02.Addr.Region)
	fmt.Printf("%#v\n", me02.Addr)
}
