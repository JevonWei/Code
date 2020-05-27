package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/JevonWei/usermanager/users"
)

func main() {

	users.Init()

	auth := flag.Bool("N", false, "No Auth")
	help := flag.Bool("h", false, "Help")
	flag.Parse()

	flag.Usage = func() {
		fmt.Println("usage flagargs [-N]")
		flag.PrintDefaults()
	}
	if *help {
		flag.Usage()
	}

	if *help || !*auth && !users.Auth() {
		return
	}

	callbacks := map[string]func(){
		"1": users.Users_Show,
		"2": users.Query,
		"3": users.User_Add,
		"4": users.User_Modify,
		"5": users.User_Delete,
		"6": users.Passwd_Modify,
		"7": func() {
			os.Exit(0)
		},
	}

	for {
		fmt.Printf(users.Info)

		if callback, ok := callbacks[users.InputString("请选择你的操作:")]; ok {
			callback()
		} else {
			fmt.Printf("你的选择有误，请重新输入")
		}

		// op := users.InputString("请选择你的操作:")

		// switch op {
		// case "1":
		// 	users.Users_Show()
		// case "2":
		// 	users.Query()
		// case "3":
		// 	users.User_Add()
		// case "4":
		// 	users.User_Modify()
		// case "5":
		// 	users.User_Delete()
		// case "6":
		// 	os.Exit(0)
		// default:
		// 	fmt.Println("你的选择无效，请重新输入")
		// }

	}
}
