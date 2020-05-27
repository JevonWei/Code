package main

import "fmt"

type Dog struct {
	name string
}

func (dog Dog) Call() {
	fmt.Printf("%s: wangwang\n", dog.name)
}

func (dog Dog) SetName(name string) {
	dog.name = name
}

func (dog *Dog) PsetName(name string) {
	dog.name = name
}

func main() {
	dog := Dog{"QQQ"}
	dog.Call()
	//dog.Name = "AAA"
	//dog.Call()

	dog.SetName("ZZZZ")
	dog.Call()

	(&dog).PsetName("ZZZZ") // 取引用
	dog.Call()              // 自动取引用，语法糖
	dog.PsetName("白")       // 自动取引用，语法糖
	dog.Call()

	pdog := &Dog{"黑"}
	pdog.Call()    // 自动解引用，语法糖
	(*pdog).Call() // 解引用
	pdog.PsetName("红")
	(pdog).Call() // 自动解引用，语法糖

}
