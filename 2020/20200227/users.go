package main

import (
	"fmt"
	"strings"
)

// 添加用户
func add(pk int, users map[string]map[string]string) {
	var (
		id   string = fmt.Sprintf("%d", pk)
		name string
		age  string
		tel  string
		addr string
	)

	fmt.Print("请输入姓名: ")
	fmt.Scan(&name)
	fmt.Print("请输入年龄: ")
	fmt.Scan(&age)
	fmt.Print("请输入电话: ")
	fmt.Scan(&tel)
	fmt.Print("请输入地址: ")
	fmt.Scan(&addr)

	users[id] = map[string]string{
		"id":   id,
		"name": name,
		"tel":  tel,
		"age":  age,
		"addr": addr,
	}
	// fmt.Println(id, name, age, tel, addr)
}

// 查询用户
func query(users map[string]map[string]string) {
	var q string
	fmt.Print("请输入查询信息:")
	fmt.Scan(&q)

	title := fmt.Sprintf("%5s|%20s|%5s|%20s|%50s", "ID", "Name", "Age", "Tel", "Addr")
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", len(title)))

	for _, user := range users {
		if q == "" || strings.Contains(user["name"], q) || strings.Contains(user["age"], q) || strings.Contains(user["tel"], q) || strings.Contains(user["addr"], q) {
			fmt.Printf("%5s|%20s|%5s|%20s|%50s", user["id"], user["name"], user["age"], user["tel"], user["addr"])
			fmt.Println()
		}
	}
}

func main() {
	// 存储用户信息
	users := make(map[string]map[string]string)
	id := 0
	fmt.Println("欢迎使用用户管理命令行系统")

END:
	for {
		var op string
		fmt.Print(`
1. 新建用户
2. 修改用户
3. 删除用户
4. 查询用户
5. 退出系统
请输入你的指令:`)

		fmt.Scan(&op)
		// fmt.Println("你输入的指令为: ", op)

		switch op {
		case "1":
			id++
			add(id, users)
		case "2":
		case "3":
		case "4":
			query(users)
		case "5":
			break END
		default:
			fmt.Println("输入的指令错误\n\n")
		}

	}
}
