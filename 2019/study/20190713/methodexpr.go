package main

import "fmt"

type Dog struct {
	name string
}

func (dog Dog) Call() {
	fmt.Printf("%s: 旺旺\n", dog.name)
}

/*
go编辑器自动生成
func (dog *Dog) Call() {
	fmt.Printf("%s: 旺旺\n", dog.name)
}
*/
func (dog *Dog) SetName(name string) {
	dog.name = name
}

func main() {
	//m1 := Dog.Call
	m1 := (*Dog).Call // 可传递指针参数
	m2 := (*Dog).SetName

	fmt.Printf("%T, %T\n", m1, m2)

	dog := Dog{"豆豆"}
	//m1(dog)
	m1(&dog)
	m2(&dog, "小黑")
	//m1(dog)
	m1(&dog)
	dog.SetName("小白")
	//m1(dog)
	m1(&dog)

	pdog := &Dog{"豆豆"}
	//m1(*pdog)
	m1(pdog)
	m2(pdog, "小黑")
	//m1(*pdog)
	m1(pdog)
	m2(pdog, "小白")
	//m1(*pdog)
	m1(pdog)

}
