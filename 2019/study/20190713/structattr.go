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
	me := User{
		ID:       1,
		Name:     "Jevon",
		Birthday: time.Now(),
		Addr:     "shanghai",
		Tel:      "1234566",
	}

	fmt.Println(me.ID, me.Name, me.Tel)
	me.Tel = "987654321"
	fmt.Printf("%#v\n", me)

	me2 := &User{
		ID:   2,
		Name: "Danran",
	}
	fmt.Printf("me2 修改前：%#v\n", me2)
	(*me2).Tel = "1222222"
	me2.Addr = "shanghai"
	fmt.Printf("me2 修改后：%#v\n", me2)

}
