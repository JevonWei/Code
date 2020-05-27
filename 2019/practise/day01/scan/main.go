package main

import "fmt"

func main() {
	var Name string
	fmt.Print("请输入你的名字:")
	fmt.Scan(&Name)
	fmt.Printf("Name: %s", Name)
}
