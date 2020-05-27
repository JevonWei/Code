package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/JevonWei/usermanager/auth"
	"github.com/JevonWei/usermanager/inits"
	"github.com/JevonWei/usermanager/inputstring"
	"github.com/JevonWei/usermanager/operate"
	"github.com/JevonWei/usermanager/usersort"
)

func main() {
	inits.Title_String()

	// 命令行输入参数-N，则跳过密码验证
	var noauth bool
	flag.BoolVar(&noauth, "N", false, "no auth")

	flag.Parse()

	if !noauth && !auth.Auth() {
		return
	}
	//END:

	// 定义map类型callbacks，返回值为函数
	callbacks := map[string]func(){
		"1": usersort.Print_sort,
		"2": operate.Query,
		"3": operate.Add,
		"4": operate.Modify,
		"5": operate.Deluser,
		"6": func() {
			os.Exit(0)
		},
	}

	for {
		fmt.Println(strings.Repeat("-", 30))
		fmt.Println(inits.Menu)

		if callback, ok := callbacks[inputstring.InputString("请输入你选择的操作:")]; ok {
			callback()
		} else {
			print("选择无效，请重新输入!!!")
		}
	}

	// for {
	// 	fmt.Println(strings.Repeat("-", 30))
	// 	fmt.Println(inits.Menu)

	// 	op := inputstring.InputString("请输入你选择的操作:")
	// 	switch op {
	// 	case "1":
	// 		usersort.Print_sort()
	// 	case "2":
	// 		operate.Query()
	// 	case "3":
	// 		operate.Add()
	// 	case "4":
	// 		operate.Modify()
	// 	case "5":
	// 		operate.Deluser()
	// 	case "6":
	// 		break END
	// 	default:
	// 		fmt.Println("选择无效，请重新输入!!!")
	// 	}
	// }

}
