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

func NewUser(id int, name string, region, street, no string) User {
	return User{
		ID:   id,
		Name: name,
		Addr: &Address{region, street, no},
	}
}

func main() {
	me := User{
		ID:   1,
		Name: "Jevon",
		Addr: &Address{"shanghai", "shuangqiao", "110"},
	}
	me2 := me
	me2.Name = "danran"
	me2.Addr.Street = "jufeng"

	fmt.Printf("%#v\n", me.Addr)
	fmt.Printf("%#v\n", me2.Addr)

	danran := NewUser(1, "dan", "shanghai", "shuangqiao", "100")
	fmt.Printf("%#v\n", danran)

}
