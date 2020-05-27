package main

import (
	"fmt"
	lgopkg "gopkg" // 别名导入

	// . "github.com/JevonWei/gopkg" // 简导入
	// _ "github.com/JevonWei/gopkg" // 只声明不引用
	"github.com/JevonWei/gopkg"
	"github.com/howeyc/gopass"
)

func main() {
	fmt.Println(gopkg.VERSION)
	fmt.Println(lgopkg.VERSION)
	lgopkg.Printname()
	fmt.Print("请输入密码：")
	if bytes, err := gopass.GetPasswd(); err == nil {
		 fmt.Println(string(bytes))
		//if string(bytes) == "123" {
		//	fmt.Println("True")
		//}
	}
}
