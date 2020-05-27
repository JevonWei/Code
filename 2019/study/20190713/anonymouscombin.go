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
	User
	Salary float64
	Name   string
}

func main() {
	var me Employee
	fmt.Printf("%#v\n", me)

	me1 := Employee{
		User: User{
			ID:   1,
			Name: "Dan",
			Addr: Address{"shanghai", "shuangqian", "1139"},
		},
		Salary: 1000,
	}
	fmt.Printf("%T, %#v\n", me1, me1)

	fmt.Println(me1.User.Addr)
	me1.User.Addr = Address{"pudong", "jinqiao", "1200"}
	fmt.Printf("%#v\n", me1.User)

	me1.Name = "JevonWei"
	me1.Addr.No = "999"
	me1.Name = "wei"
	fmt.Printf("%#v\n", me1)

}
