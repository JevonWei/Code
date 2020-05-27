package main

import "fmt"

type User struct {
	id   int
	name string
}

func (user User) GetID() int {
	return user.id
}

func (user User) GetName() string {
	return user.name
}

func (user *User) SetID(id int) {
	user.id = id
}

func (user *User) SetName(name string) {
	user.name = name
}

type Employee struct {
	User
	Salary float64
	name   string
}

func (employee Employee) GetName() string {
	return employee.name
}

func (employee *Employee) SetName(name string) {
	employee.name = name
}

func main() {
	var me Employee = Employee{
		User:   User{1, "JevonWei"},
		Salary: 1000,
		name:   "kaka",
	}

	fmt.Println(me.User.GetName())
	fmt.Println(me.GetName())
	fmt.Printf("%#v\n", me)
	me.SetName("Danran")
	fmt.Println(me.GetName())
	fmt.Printf("%#v\n", me)
}
