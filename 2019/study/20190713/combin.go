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

func main() {
	var me01 User
	fmt.Printf("%#v\n", me01)

	addr := Address{"shanghai", "pudong", "02"}
	// User结构体中嵌套结构体Address
	me02 := User{
		ID:   1,
		Name: "Jevonwei",
		Addr: addr,
	}
	fmt.Printf("%#v\n", me02)

	me03 := User{
		ID:   2,
		Name: "Danran",
		Addr: Address{
			Region: "shanghai",
			Street: "shuangqiao",
			No:     "1139",
		},
	}
	fmt.Printf("%#v\n", me03)
	me03.Addr = Address{
		Region: "asdf",
		Street: "tang",
		No:     "11",
	}
	fmt.Printf("%#v\n", me03.Addr)
	me03.Addr.No = "123"
	fmt.Printf("%#v\n", me03.Addr)
}
