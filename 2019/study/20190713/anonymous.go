package main

import "fmt"

func main() {
	var me struct {
		ID   int
		Name string
	}

	fmt.Printf("%T\n", me)
	fmt.Printf("%#v\n", me)

	me.Name = "danran"
	me.ID = 1
	fmt.Printf("%#v\n", me)
	fmt.Println(me.Name)

	// 匿名结构体字面量初始化结构体
	me2 := struct {
		ID   int
		Name string
	}{1, "Jevon"}

	fmt.Printf("%#v\n", me2)

}
