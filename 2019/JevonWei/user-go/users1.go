package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/howeyc/gopass"
)

const (
	maxAuth  = 3
	passWord = "danran"
)

// print打印函数
func print(s string) {
	fmt.Println(s)
}

// 从命令行输入函数
func inputString(s string) string {
	var input string
	fmt.Print(s)
	fmt.Scan(&input)
	return strings.TrimSpace(input)
}

// 认证函数，从命令行输入密码，并进行验证
// 通过返回值验证是否成功
func auth() bool {
	//var input string
	for i := 0; i < maxAuth; i++ {
		fmt.Println("请输入本系统登录密码:")
		//fmt.Scan(&input)
		//if passWord == input {
		// 	return true
		// } else {
		// 	fmt.Println("密码输入错误，请重新输入...")
		// 	fmt.Println(strings.Repeat("-", 30))
		// }

		// 输入密码不回显
		if bytes, err := gopass.GetPasswd(); err == nil {
			if string(bytes) == passWord {
				return true
			} else {
				fmt.Println("密码输入错误，请重新输入...")
				fmt.Println(strings.Repeat("-", 30))
			}
		}
	}
	return false
}

// 显示选择的用户
func listUser(pk int, user map[string]string) {
	print("================================")
	fmt.Println("ID:", pk)
	fmt.Println("Name:", user["name"])
	fmt.Println("出生日期:", user["birthday"])
	fmt.Println("联系方式:", user["tel"])
	fmt.Println("地址:", user["addr"])
}

// 打印所有的用户信息
func printUser(users map[int]map[string]string) {
	for id, user := range users {
		listUser(id, user)
	}
	print("================================")
}

// 查询函数
func query(users map[int]map[string]string) {
	q := inputString("请输入查询的信息:")

	title := fmt.Sprintf("%-5s|%-10s|%-5s|%-10s|%-15s", "ID", "Name", "birthday", "Tel", "Addr")
	print(title)
	print((strings.Repeat("-", len(title))))
	for k, user := range users {
		if strings.Contains(user["name"], q) || strings.Contains(user["tel"], q) || strings.Contains(user["addr"], q) {
			fmt.Printf("%-5d|%-10s|%-5s|%-10s|%-15s\n\n", k, user["name"], user["birthday"], user["tel"], user["addr"])

		}
	}
}

// 获取users中的用户的最大ID
func getId(users map[int]map[string]string) int {
	var id int
	for k := range users {
		if id < k {
			id = k
		}
	}
	return id + 1
}

// 用户信息
func inputUser() map[string]string {
	user := map[string]string{}

	user["name"] = inputString("请输入名字:")
	user["birthday"] = inputString("请输入出生日期(2019-07-07):")
	user["tel"] = inputString("请输入联系方式:")
	user["addr"] = inputString("请输入地址:")
	print("*******************************")
	return user
}

// 添加函数
func add(users map[int]map[string]string) {
	id := getId(users)

	// 调用用户函数，新增用户
	user := inputUser()

	users[id] = user
	fmt.Printf("ID为%d的用户已添加\n", id)

	// name := inputString("请输入名字: ")
	// birthday := inputString("请输入出生日期(2019-07-07)：")
	// tel := inputString("请输入联系方式: ")
	// addr := inputString("请输入地址: ")

	// users[id] = {
	// 	"name" := name,
	// 	"birthday" := birthday,
	// 	"tel" := tel,
	// 	"addr" := addr,
	// }
}

// 修改函数
func modify(users map[int]map[string]string) {
	printUser(users)

	idString := inputString("请输入修改用户ID:")

	if id, err := strconv.Atoi(idString); err == nil {
		if user, ok := users[id]; ok {
			fmt.Println("")
			print("将要修改的用户信息为:")
			listUser(id, user)
			input := inputString("是否确定修改(Y/N)?: ")
			if input == "Y" || input == "y" {
				user := inputUser()
				users[id] = user
				fmt.Printf("ID为%d的用户已修改\n", id)
			}
		} else {
			print("输入的用户ID不存在")
		}
	} else {
		print("输入的ID不正确")
	}
}

// 删除用户
func delUser(users map[int]map[string]string) {
	printUser(users)

	idString := inputString("请输入删除用户ID:")
	if id, err := strconv.Atoi(idString); err == nil {
		if user, ok := users[id]; ok {
			print("将要删除的用户信息为:")
			listUser(id, user)
			input := inputString("是否确定删除(Y/N)?")
			if input == "Y" || input == "y" {
				delete(users, id)
				fmt.Printf("ID为%d的用户已删除\n", id)
			}
		} else {
			print("输入的用户ID不存在")
		}
	} else {
		print("输入的ID不正确")
	}
}

func main() {
	if !auth() {
		fmt.Printf("密码输入%d次错误，程序退出\n", maxAuth)
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
		"1": query,
		"2": add,
		"3": modify,
		"4": delUser,
		"5": func(users map[int]map[string]string) {
			os.Exit(0)
		},
	}
	//END:
	for {
		print("")
		print(strings.Repeat("-", len(head)))
		print(menu)

		if callback, ok := callbacks[inputString("请输入你选择的操作:")]; ok {
			callback(users)
		} else {
			print("选择无效，请重新输入!!!")
		}

		// 示范二
		//op := inputString("请输入你选择的操作:")
		//if callback, ok := callbacks[op]; ok {
		//	callback(users)
		//} else if op == "5" {
		//	break
		//} else {
		//	print("选择无效，请重新输入!!!")
		//}

		// 示范一
		// op := inputString("请输入你选择的操作:")
		// switch op {
		// case "1":
		// 	query(users)
		// case "2":
		// 	add(users)
		// case "3":
		// 	modify(users)
		// case "4":
		// 	delUser(users)
		// case "5":
		// 	break END
		// default:
		// 	print("选择无效，请重新输入!!!")
		// }
	}

}
