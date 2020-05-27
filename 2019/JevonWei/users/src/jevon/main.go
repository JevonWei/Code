package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/JevonWei/user/auth"
	"github.com/JevonWei/user/input"
	"github.com/JevonWei/user/operate"
)

func print(s string) {
	fmt.Println(s)
}

func main() {
	if !auth.Auth() {
		return
	}

	menu := `1. 查询
2. 添加
3. 修改
4. 删除
5. 退出
*******************************`

	head := "欢迎进入JevonWei的用户管理系统"
	print(head)

	users := map[int]map[string]string{}

	callbacks := map[string]func(map[int]map[string]string){
		"1": operate.Query,
		"2": operate.Add,
		"3": operate.Modify,
		"4": operate.DelUser,
		"5": func(users map[int]map[string]string) {
			os.Exit(0)
		},
	}
	//END:
	for {
		print("")
		print(strings.Repeat("-", len(head)))
		print(menu)

		if callback, ok := callbacks[input.InputString("请输入你选择的操作:")]; ok {
			callback(users)
		} else {
			print("选择无效，请重新输入!!!")
		}
	}
}
