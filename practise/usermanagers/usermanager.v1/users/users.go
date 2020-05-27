package users

import (
	"fmt"
	"strconv"
	"strings"
)

// 系统主函数，用户Add，Query，Modify，Delete

// type User struct {
// 	ID       int
// 	Name     string
// 	Birthday time.Time
// 	Addr     string
// 	Tel      string
// 	Desc     string
// }

// var Users map[int]User = map[int]User{}

func (U User) Show(user User) {
	fmt.Println("================================")
	fmt.Println("ID:", user.ID)
	fmt.Println("Name:", user.Name)
	fmt.Println("出生日期:", user.Birthday.Format("2006/01/02"))
	fmt.Println("联系方式:", user.Tel)
	fmt.Println("地址:", user.Addr)
	fmt.Println("描述:", user.Desc)
}

func User_Add() {
	id := GetId()
	InputUser(id)

}

func Users_Show() {
	// fmt.Println(Users)
	// title := fmt.Sprintf("%-5s|%-10s|%-15s|%-10s|%-15s|%-15s", "ID", "Name", "Birthday", "Tel", "Addr", "Desc")
	// fmt.Println(title)
	// fmt.Println((strings.Repeat("-", len(title))))

	// 打印用户显示的Title信息
	Print_Title()

	// Users_Sort()将用户按照选择的字段排序
	for _, user := range Users_Sort() {
		fmt.Printf("%-5d|%-10s|%-15s|%-10s|%-15s|%-15s\n", user.ID, user.Name, user.Birthday.Format("2006/01/02"), user.Tel, user.Addr, user.Desc)
	}
}

func Query() {
	q := InputString("请输入要查询的信息: ")
	User_Sorted := Users_Sort()
	Print_Title()

	for _, user := range User_Sorted {
		if strings.Contains(user.Name, q) || strings.Contains(user.Addr, q) || strings.Contains(user.Desc, q) || strings.Contains(user.Tel, q) {
			fmt.Printf("%-5d|%-10s|%-15s|%-10s|%-15s|%-15s\n", user.ID, user.Name, user.Birthday.Format("2006/01/02"), user.Tel, user.Addr, user.Desc)
		} else {
			fmt.Println("查询的用户不存在")
		}
	}
}

func User_Modify() {
	idString := InputString("请输入要修改用户的ID:")
	if id, err := strconv.Atoi(idString); err == nil {
		if user, ok := Users[id]; ok {
			fmt.Printf("\n将要修改的用户信息为:\n")

			User.Show(user, user)

			confirm := InputString("是否确认修改(Y/N)?: ")
			if confirm == "Y" || confirm == "y" {
				InputUser(id)
				fmt.Printf("ID为%d的用户已修改\n", id)
			}
		} else {
			fmt.Println("输入的用户ID不存在.")
		}
	} else {
		fmt.Println("ID输入错误.")
	}
}

func User_Delete() {
	idString := InputString("请输入要删除用户的ID:")
	if id, err := strconv.Atoi(idString); err == nil {
		if user, ok := Users[id]; ok {

			fmt.Printf("\n将要删除的用户信息为:\n")
			User.Show(user, user)

			confirm := InputString("是否确认删除(Y/N)?: ")
			if confirm == "Y" || confirm == "y" {
				delete(Users, id)
				fmt.Printf("ID为%d的用户已删除\n", id)
			}
		} else {
			fmt.Println("输入的用户ID不存在.")
		}
	} else {
		fmt.Println("ID输入错误.")
	}
}
