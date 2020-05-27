package users

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

// 系统辅助程序函数

// 将Birthday字符串装换为time.Time类型
func Birthday_time(s string) (time.Time, error) {

	T, Err := time.Parse("2006-01-02", s)

	return T, Err

}

// 获取新添加用户的ID
func GetId() int {
	var Id int
	for k := range Users {
		if k > Id {
			Id = k
		}
	}
	return Id + 1
}

// 从键盘输入用户的信息
func InputUser(num int) {
	var User User = User{}

	name := InputString("请输入名字:")
	if name == "" {
		fmt.Println("输入的Name不能为空")
		goto END
	}

	for _, user := range Users {
		if num != user.ID && name == user.Name {
			fmt.Printf("用户名%s已存在,不能新增/修改\n", name)
			goto END
		}
	}

	User.Name = name
	User.ID = num
	for {
		birthday_time, err := Birthday_time(InputString("请输入出生日期(2019-07-07):"))
		if err == nil {
			User.Birthday = birthday_time
			break
		} else {
			fmt.Println(errors.New("输入的格式错误!"))
		}
	}
	User.Tel = InputString("请输入联系方式:")
	User.Addr = InputString("请输入地址:")
	User.Desc = InputString("请输入用户的描述信息:")

	fmt.Println("*******************************")
	// fmt.Printf("ID为%d的用户已添加/修改\n", num)

	Users[num] = User
END:
}

func Users_Sort() []User {
	// 定义User的空数组
	Users_Array := []User{}

	fmt.Println(Sort_Menu)
	sort_value := InputString("请选择要排序的字段: ")

	// 将所有的用户信息保存在数组中
	for _, user := range Users {
		Users_Array = append(Users_Array, user)
	}

	// 调用sort.Slice()函数，根据选择的字段排序用户
	switch sort_value {
	case "1":
		sort.Slice(Users_Array, func(i, j int) bool {
			return Users_Array[i].ID < Users_Array[j].ID
		})
	case "2":
		sort.Slice(Users_Array, func(i, j int) bool {
			return Users_Array[i].Name < Users_Array[j].Name
		})
	case "3":
		sort.Slice(Users_Array, func(i, j int) bool {
			return Users_Array[i].Birthday.Format("2006/01/02") < Users_Array[j].Birthday.Format("2006/01/02") // 将time类型的Birthday值转换为字符串排序
		})
	case "4":
		sort.Slice(Users_Array, func(i, j int) bool {
			return Users_Array[i].Addr < Users_Array[j].Addr
		})
	case "5":
		sort.Slice(Users_Array, func(i, j int) bool {
			return Users_Array[i].Tel < Users_Array[j].Tel
		})
	case "6":
		sort.Slice(Users_Array, func(i, j int) bool {
			return Users_Array[i].Desc < Users_Array[j].Desc
		})
	case "7":
		break
	}
	return Users_Array
}
