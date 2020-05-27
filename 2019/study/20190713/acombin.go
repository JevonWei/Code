package main

import "fmt"

type Address struct {
	Region string
	Street string
	No     string
}

type Company struct {
	ID     int
	Name   string
	Addr   Address
	Salary float64
}

type User struct {
	ID   int
	Name string
	Addr Address
}

type Employee struct {
	User
	Company
	Salary float64
	Name   string
}

func main() {
	var me Employee
	fmt.Printf("%T, %#v\n", me, me)

	me.Company.Name = "AAA"
	me.User.Name = "BB"
	fmt.Println(me.Company.Name)
	fmt.Println(me.User.Name)

	fmt.Printf("%#v\n", me)
}
