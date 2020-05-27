package main

import (
	"fmt"
	"time"
)

type User struct {
	ID       int
	Name     string
	Birthday time.Time
	Addr     string
	Tel      string
	Remark   string
}

func main() {
	var me User
	fmt.Printf("%T\n", me)
	fmt.Printf("%#v\n", me)

	var me2 User = User{1, "JevonWei", time.Now(), "shanghai", "1234566", "Jevon"}
	fmt.Printf("me2 值为：%#v\n", me2)

	var me3 User = User{}
	fmt.Printf("me3 值为：%#v\n", me3)

	var me4 User = User{ID: 1,
		Name:     "Jevon",
		Birthday: time.Now(),
		Addr:     "shanghai",
		Tel:      "1234566",
	}
	fmt.Printf("me4 值为: %#v\n", me4)

	var pointer *User
	fmt.Printf("%#v\n", pointer)

	var pointer2 *User = &me4
	fmt.Printf("%#v\n", pointer2)

	var pointer3 *User = &User{}
	fmt.Printf("%#v\n", pointer3)

	// new() 函数创建的是指针类型
	var pointer4 *User = new(User)
	fmt.Printf("%#v\n", pointer4)

	var pointer5 *int = new(int)
	fmt.Printf("%#v\n", pointer5)

}
