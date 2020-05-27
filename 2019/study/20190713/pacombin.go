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

type Employee struct {
	*User
	Salary float64
	Name   string
}

func main() {
	var me Employee
	fmt.Printf("%#v\n", me)

	me1 := Employee{
		User: &User{
			ID:   1,
			Name: "danran",
			Addr: Address{"shanghai", "shuangqiao", "1109"},
		},
		Salary: 100,
		Name:   "Jevon",
	}
	fmt.Printf("%#v\n", me1)

	me1.User.Name = "wei"
	fmt.Println(me1.Name)
	fmt.Println(me1.User.Name)
	fmt.Println(me1.Addr)
}
