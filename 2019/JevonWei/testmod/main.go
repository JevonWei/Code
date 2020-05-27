package main

import (
	"fmt"

	//"github.com/JevonWei/gopkg"
	"github.com/JevonWei/testmod/gopkg"
	"github.com/astaxie/beego"
	"github.com/howeyc/gopass"
)

func main() {
	gopass.GetPasswd()
	fmt.Println(gopkg.Version)
	beego.Run()
}
