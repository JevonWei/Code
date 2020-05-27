package main

import (
	"fmt"
	"strings"
)

func Input_string(input_string *string, message string) {
	fmt.Print(message)
	fmt.Scan(input_string)
}

func Add(pk int, users map[string]map[string]string) {
	var (
		id   string = fmt.Sprintf("%d", pk)
		name string
		age  string
		tel  string
		addr string
	)
	Input_string(&name, "请输入Name: ")
	Input_string(&age, "请输入Age: ")
	Input_string(&tel, "请输入Tel: ")
	Input_string(&addr, "请输入Addr: ")

	users[id] = map[string]string{
		"id":   id,
		"name": name,
		"age":  age,
		"tel":  tel,
		"addr": addr,
	}
}

func query(users map[string]map[string]string) {
	var q string
	Input_string(&q, "请输入要查询的信息: ")

	title := fmt.Sprintf("%-2s|%-10s|%-5s|%-10s|%-20s", "ID", "Name", "Age", "Tel", "Addr")
	fmt.Println(title)
	fmt.Println(strings.Repeat("*", len(title)))

	for _, user := range users {
		if q == "" || strings.Contains(user["name"], q) || strings.Contains(user["tel"], q) || strings.Contains(user["addr"], q) || strings.Contains(user["age"], q) {
			fmt.Printf("%-2s|%-10s|%-5s|%-10s|%-20s", user["id"], user["name"], user["age"], user["tel"], user["addr"])
			fmt.Println()
		}
	}
}

func modify(users map[string]map[string]string) {
	var pk string
	var op string
	var info string
	var IsVerify string

	Input_string(&pk, "请输入要修改的用户ID: ")

	if _, ok := users[pk]; ok {
		message := `
请选择要修改的字段序列化:
1. Name
2. Age
3. Tel
4. Addr
请选择要修改的内容: `

		Input_string(&op, message)

		IsVerify_message := fmt.Sprintf("是否确认修改ID为[%s]的用户的信息(Y?N)", pk)
		Input_string(&IsVerify, IsVerify_message)
		if IsVerify == "Y" || IsVerify == "y" {
			switch op {
			case "1":
				Input_string(&info, "请输入Name: ")
				users[pk]["name"] = info
				fmt.Printf("用户ID为[%s]的Name已修改\n", pk)
			case "2":
				Input_string(&info, "请输入Age: ")
				users[pk]["age"] = info
				fmt.Printf("用户ID为[%s]的age已修改\n", pk)
			case "3":
				Input_string(&info, "请输入Tel: ")
				users[pk]["tel"] = info
				fmt.Printf("用户ID为[%s]的tel已修改\n", pk)
			case "4":
				Input_string(&info, "请输入Addr: ")
				users[pk]["addr"] = info
				fmt.Printf("用户ID为[%s]的addr已修改\n", pk)
			default:
				fmt.Println("输入的信息有误")

			}
		} else if IsVerify == "N" || IsVerify == "n" {
			fmt.Printf("您取消了用户ID为[%s]的修改\n", users[pk])
		}
	} else {
		fmt.Println("输入的用户ID不存在，请重新输入!!!")
	}
}

func remove(users map[string]map[string]string) {
	var pk string
	var Is_Delete string
	Input_string(&pk, "输入要删除的用户ID: ")

	if _, ok := users[pk]; ok {
		message := fmt.Sprintf("是否要删除ID为%s的用户(Y/N)", pk)
		Input_string(&Is_Delete, message)
		if Is_Delete == "Y" || Is_Delete == "y" {
			fmt.Println(users[pk])
			delete(users, pk)
		} else if Is_Delete == "N" || Is_Delete == "n" {
			fmt.Printf("您取消了ID为%s的删除操作\n", pk)
		}
	} else {
		fmt.Println("删除的用户不存在，请重新输入...")
	}

}

func main() {
	password := "123123"
	var passwd string
	var reset string

	defer func() {
		fmt.Println("欢迎使用本系统")
	}()

	for i := 1; i <= 3; i++ {
		Input_string(&passwd, "请输入系统密码: ")
		if passwd == password {
			fmt.Println("欢迎进入用户管理命令行系统")
			id := 0
			users := make(map[string]map[string]string)

		END:
			for {
				var op string
				message := `
1. 新建用户
2. 修改用户
3. 删除用户
4. 查询用户
5. 退出系统
请输入你的指令:`
				Input_string(&op, message)

				switch op {
				case "1":
					id++
					Add(id, users)
				case "2":
					modify(users)
				case "3":
					remove(users)
				case "4":
					query(users)
				case "5":
					break END
				default:
					fmt.Println("输入的指令有误")
				}
			}
		} else {
			fmt.Println("密码输入错误，请重新输入...")
		}

		if i == 3 && passwd != password {
			Input_string(&reset, "密码输入错误，系统即将退出，是否重新开始输入密码(Y?N)!")
			if reset == "Y" || reset == "y" {
				i = 0
			}
		}

	}

}
