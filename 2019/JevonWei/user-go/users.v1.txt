package users

import (
	"bufio"
	"crypto/md5"
	"encoding/csv"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/howeyc/gopass"
)

// 定义常量maxAuth，为输入密码的最多次数
// 定义常量password，为密码的MD5加密值
const (
	maxAuth    = 3
	passwdfile = "passwd.gob"
	user_file  = "user.gob"
)

// 定义变量Menu为操作的可选项
var Menu = `1. 显示
2. 查询
3. 添加
4. 修改
5. 删除
6. 退出
*******************************`

// 定义变量显示所有的Users结构体类型的所有参数
var Sort_menu = `1. ID
2. Name
3. Birthday
4. Addr
5. Tel
6. Desc
*******************************`

// 在用户系统登录前，显示提示信息
func Title_String() {
	fmt.Println("JevonWei用户系统密码为:danran")
	fmt.Println("")

	// strings.Repeat() 显示特定字符多少次
	fmt.Println(strings.Repeat("*", 30))
	Head := "欢迎进入JevonWei的用户管理系统"
	fmt.Println(Head)
}

// 定义用户结构体类型Users

type Users struct {
	ID       int
	Name     string
	Birthday time.Time
	Addr     string
	Tel      string
	Desc     string
}

// 定义用户变量

var User map[int]Users = map[int]Users{}

func (u Users) String() string {
	return fmt.Sprintf("ID: %d\n名字: %s\n出生日期: %s\n联系方式: %s\n联系地址: %s\n备注: %s", u.ID, u.Name, u.Birthday.Format("2006/01/02"), u.Addr, u.Tel, u.Desc)
}

// 定义从键盘输入函数，并返回输入的值
func InputString(s string) string {
	var in string
	fmt.Print(s)
	fmt.Scan(&in)
	return strings.TrimSpace(in)
}

func Encodepasswd() {
	bytes, _ := gopass.GetPasswd()

	var password = fmt.Sprintf("%x", md5.Sum(bytes))

	file, err := os.Create("passwdfile")
	if err == nil {
		defer file.Close()

		encoder := gob.NewEncoder(file)
		encoder.Encode(password)
		// 将输入的密码写入csv文件中
		csvwpasswd(password)
	}
}

func Decodepasswd() string {
	var password string

	file, err := os.Open("passwdfile")
	if err == nil {
		defer file.Close()

		decoder := gob.NewDecoder(file)
		decoder.Decode(&password)

		return password
	}
	return ""
}

func csvwpasswd(password string) {

	file, err := os.Create("passwdfile.csv")

	if err == nil {
		defer file.Close()

		wirter := csv.NewWriter(file)
		wirter.Write([]string{password})

		wirter.Flush()
	}
}

func csvrpasswd() []string {
	file, err := os.Open("passwdfile.csv")
	if err == nil {
		defer file.Close()

		reader := csv.NewReader(file)

		for {
			line, err := reader.Read()
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
				return []string{}
			} else {
				return line
			}
			return []string{}
		}

	}
	return []string{}
}

func Userencode(users map[int]Users) {
	file, err := os.Create(user_file)
	if err != nil {
		fmt.Println("文件打开失败")
		fmt.Println(err)
	} else {
		defer file.Close()
		gobencode := gob.NewEncoder(file)
		gobencode.Encode(users)
	}
}

// 反序列化
func Userdecode() map[int]Users {
	users := map[int]Users{}
	file, err := os.Open(user_file)
	if err == nil {
		defer file.Close()

		decoder := gob.NewDecoder(file)
		decoder.Decode(&users)

		return users
	}
	return nil

}

//func CsvwUser(num, ID, Name, Birthday, Addr, Tel, Desc string) {
// func CsvwUser(users map[int]Users) {
// 	file, err := os.OpenFile(user_file, os.O_RDWR|os.O_CREATE, 0666)
// 	//file, err := os.Open(user_file)
// 	if err != nil {
// 		fmt.Println("文件打开失败")
// 		fmt.Println(err)
// 	}

// 	defer file.Close()
// 	file.Seek(0, io.SeekEnd)

// 	writer := csv.NewWriter(file)
// 	writer.Write(users)
// 	writer.Flush()

// }

// func CsvrUser() [][]string {
// 	var lines [][]string = [][]string{}
// 	file, err := os.Open(user_file)
// 	if err != nil {
// 		fmt.Println("文件打开失败")
// 		return [][]string{}
// 	}

// 	defer file.Close()
// 	reader := csv.NewReader(file)
// 	for {
// 		line, err := reader.Read()
// 		if err != nil {
// 			if err != io.EOF {
// 				fmt.Println(err)
// 				return [][]string{}
// 			}
// 			break
// 		} else {
// 			lines = append(lines, line)
// 		}

// 	}
// 	return lines
// }

// func stdin() string {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	return scanner.Text()
// }

func stdin() string {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	return str
}

// 认证函数
func Auth() bool {
	// 判断密码是否存在，若不存在，则设置密码
	file, err := os.Open("passwdfile")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("密码不存在，请设置密码")
			Encodepasswd()
			return true
		}
	} else {
		file.Close()
		for i := 0; i < maxAuth; i++ {
			fmt.Print("请输入JevonWei用户系统密码: ")
			// 将输入的密码隐形显示
			bytes, _ := gopass.GetPasswd()

			// 如果输入密码的MD5值等于password，返回True
			if Decodepasswd() == fmt.Sprintf("%x", md5.Sum(bytes)) {
				//os.Stdout.Write([]byte("是否要修改密码(Y/N)："))
				fmt.Printf("是否要修改密码(Y/N)：")
				if stdin() == "y" || stdin() == "Y" {
					os.Stdout.Write([]byte("请输入原始密码："))
					//fmt.Println(csvreader()[0])
					bytes, _ := gopass.GetPasswd()
					if fmt.Sprintf("%x", md5.Sum(bytes)) == csvrpasswd()[0] {
						Encodepasswd()
						return true
					} else {
						os.Stdout.Write([]byte("密码输入错误："))
						fmt.Println("")
						return false
					}
				} else {
					return true
				}
			} else {
				fmt.Println("密码错误")
			}
		}
	}

	fmt.Printf("密码输入%d次错误，程序退出\n", maxAuth)
	return false
}

// 将输入的字符串数据转换为时间类型，格式为2006-01-02，返回时间类型的值和错误信息
func Birthday_time(s string) (time.Time, error) {

	T, Err := time.Parse("2006-01-02", s)

	return T, Err

}

// 获取用户的最大ID，且返回ID+1
func GetId() int {
	Id := 0
	users := Userdecode()

	if len(users) == 0 {
		return 1
	}

	for k := range users {
		if Id < k {
			Id = k
		}
	}
	return Id + 1

}

// 定义用户输入函数，将从键盘输入的每个值对应传入Users结构体的元素中
func Inputuser(num int) {
	// 定义结构体类型Users的变量User_input
	var User_input Users = Users{}

	// 从键盘输入name(用户名)
	name := InputString("请输入名字:")
	// fmt.Printf("请输入Name:")
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// name := scanner.Text()

	// 若name为空，则不赋值，直接返回
	if name == "" {
		fmt.Println("输入的Name不能为空")
		goto END
	}

	// 若输入的name，在系统中已存在，则提示用户已存在，不能新建用户
	for _, user := range Userdecode() {
		if name == user.Name && user.ID != num {
			fmt.Printf("用户名%s已存在,不能新增/修改\n", name)
			goto END
		}
	}

	// 将输入的name，赋值给Users结构体的Name
	User_input.Name = name
	// 将User映射的key值，赋值给结构体Users的ID
	User_input.ID = num

	// 将输入的字符串类型的Birthday转换为时间类型
	// 判断输入的Birthday格式是否正确，若输入格式有误，则打印提示信息，并重新输入
	for {
		birthday_time, err := Birthday_time(InputString("请输入出生日期(2019-07-07):"))
		if err == nil {
			User_input.Birthday = birthday_time
			break
		} else {
			fmt.Println(errors.New("请输入正确认格式"))
		}
	}

	// 将输入的其他值依次赋值为Tel，Addr，Desc,并返回用户信息

	//User_input.Birthday = inputstring.InputString("请输入出生日期(2019-07-07):")
	User_input.Tel = InputString("请输入联系方式:")
	User_input.Addr = InputString("请输入地址:")
	User_input.Desc = InputString("请输入描述信息:")
	fmt.Println("*******************************")
	fmt.Printf("ID为%d的用户已添加/修改\n", num)

	//CsvwUser(strconv.Itoa(num), strconv.Itoa(User_input.ID), User_input.Name, User_input.Birthday.Format("2006/01/02"), User_input.Tel, User_input.Addr, User_input.Desc)
	User[num] = User_input
	Userencode(User)

END:
}

// 定义Listuser函数，打印用户系统中所有的用户
func Listuser() {

	//title := fmt.Sprintf("%-5s|%-10s|%-15s|%-10s|%-15s|%-15s", "ID", "Name", "Birthday", "Tel", "Addr", "Desc")
	//fmt.Println(title)
	//fmt.Println((strings.Repeat("-", len(title))))

	// 遍历所有的用户，并打印
	for _, user := range Userdecode() {
		fmt.Printf("%-5d|%-10s|%-15s|%-10s|%-15s|%-15s\n", user.ID, user.Name, user.Birthday.Format("2006/01/02"), user.Tel, user.Addr, user.Desc)
		//fmt.Printf("%-5s|%-10s|%-15s|%-10s|%-15s|%-15s\n", user[1], user[2], user[3], user[4], user[5], user[6])
	}
}

// 添加函数
func Add() {
	id := GetId()

	// 调用用户函数，新增用户

	Inputuser(id)

}

// 删除用户
func Deluser() {
	// 调用函数，显示系统中的所有用户
	Listuser()
	users := Userdecode()

	// 按照输入的用户ID，删除用户
	idString := InputString("请输入删除用户ID:")
	if id, err := strconv.Atoi(idString); err == nil {
		if user, ok := users[id]; ok {
			fmt.Println("将要删除的用户信息为:")
			fmt.Println("================================")
			fmt.Println(user)
			// fmt.Println("ID:", user.ID)
			// fmt.Println("Name:", user.Name)
			// fmt.Println("出生日期:", user.Birthday.Format("2006/01/02"))
			// fmt.Println("联系方式:", user.Tel)
			// fmt.Println("地址:", user.Addr)
			// fmt.Println("描述:", user.Desc)

			// 确认是否删除用户
			in := InputString("是否确定删除(Y/N)?")
			if in == "Y" || in == "y" {
				delete(users, id)
				Userencode(users)
				fmt.Printf("ID为%d的用户已删除\n", id)
			}
		} else {
			fmt.Println("输入的用户ID不存在")
		}
	} else {
		fmt.Println("输入的ID不正确")
	}
}

// 修改函数
func Modify() {
	// 调用函数，显示系统中的所有用户
	Listuser()
	users := Userdecode()
	// 根据输入的用户ID，修改用户信息
	idString := InputString("请输入修改用户ID:")

	// 将输入的字符串类型的值转换为int类型
	if id, err := strconv.Atoi(idString); err == nil {
		if user, ok := users[id]; ok {
			fmt.Println("")
			fmt.Println("将要修改的用户信息为:")
			fmt.Println("================================")
			fmt.Println(user)
			// fmt.Println("ID:", user.ID)
			// fmt.Println("Name:", user.Name)
			// fmt.Println("出生日期:", user.Birthday.Format("2006/01/02"))
			// fmt.Println("联系方式:", user.Tel)
			// fmt.Println("地址:", user.Addr)
			// fmt.Println("描述:", user.Desc)

			in := InputString("是否确定修改(Y/N)?: ")
			if in == "Y" || in == "y" {
				Inputuser(id)

			}
		} else {
			fmt.Println("输入的用户ID不存在")
		}
	} else {
		fmt.Println("输入的ID不正确")
	}
}

// // 修改函数
// func Modify() {
// 	// 调用函数，显示系统中的所有用户
// 	Listuser()

// 	// 根据输入的用户ID，修改用户信息
// 	idString := InputString("请输入修改用户ID:")

// 	// 将输入的字符串类型的值转换为int类型
// 	//if id, err := strconv.Atoi(idString); err == nil {
// 	//if idString ==  CsvrUser()[1:][1] {
// 	for _, user := range CsvrUser()[1:] {
// 		if user[1] == idString {
// 			fmt.Println("")
// 			fmt.Println("将要修改的用户信息为:")
// 			fmt.Println("================================")
// 			fmt.Println(user)
// 			// fmt.Println("ID:", user.ID)
// 			// fmt.Println("Name:", user.Name)
// 			// fmt.Println("出生日期:", user.Birthday.Format("2006/01/02"))
// 			// fmt.Println("联系方式:", user.Tel)
// 			// fmt.Println("地址:", user.Addr)
// 			// fmt.Println("描述:", user.Desc)
// 			in := InputString("是否确定修改(Y/N)?: ")
// 			if in == "Y" || in == "y" {
// 				if id, err := strconv.Atoi(idString); err == nil {
// 					Inputuser(id)
// 				}
// 			}
// 		}
// 	}
// }

// 查询函数
func Query() {
	q := InputString("请输入查询的信息:")

	title := fmt.Sprintf("%-5s|%-10s|%-15s|%-10s|%-15s|%-15s", "ID", "Name", "Birthday", "Tel", "Addr", "Desc")
	fmt.Println(title)
	fmt.Println((strings.Repeat("-", len(title))))
	//fmt.Println(CsvrUser())
	// 若输入的查询信息包含在Nmae、Desc、Addr任意一个参数中，则返回用户信息
	for _, user := range Userdecode() {
		//fmt.Println(user)
		//fmt.Printf("%-5d|%-10s|%-15s|%-10s|%-15s|%-15s\n", user.ID, user.Name, user.Birthday.Format("2006/01/02"), user.Tel, user.Addr, user.Desc)
		if strings.Contains(user.Name, q) || strings.Contains(user.Addr, q) || strings.Contains(user.Desc, q) {
			//fmt.Printf("%-5s|%-10s|%-15s|%-10s|%-15s|%-15s\n", user[1], user[2], user[3], user[4], user[5], user[6])
			fmt.Printf("%-5d|%-10s|%-15s|%-10s|%-15s|%-15s\n", user.ID, user.Name, user.Birthday.Format("2006/01/02"), user.Tel, user.Addr, user.Desc)
		}

	}

}

// 定义排序函数 ，将用户按照指定的参数排序，并返回用户数组

func User_sort() []Users {

	// 定义空数组，存储用户信息
	Users_array := []Users{}

	// 输入按照哪个参数排序，并将输入的参数赋值给变量o
	o := InputString("请输入需要排序的键值:")

	//将系统中的用户保存在Users_array数组中
	for _, user := range Userdecode() {
		Users_array = append(Users_array, user)
	}

	// 调用sort.Slice()函数，根据输入的排序参数，排序系统的所有用户
	switch o {
	case "1":
		sort.Slice(Users_array, func(i, j int) bool {
			return Users_array[i].ID < Users_array[j].ID
		})
	case "2":
		sort.Slice(Users_array, func(i, j int) bool {
			return Users_array[i].Name < Users_array[j].Name
		})
	case "3":
		sort.Slice(Users_array, func(i, j int) bool {
			return Users_array[i].Birthday.Format("2006/01/02") < Users_array[j].Birthday.Format("2006/01/02") // 将time类型的Birthday值转换为字符串排序
		})
	case "4":
		sort.Slice(Users_array, func(i, j int) bool {
			return Users_array[i].Addr < Users_array[j].Addr
		})
	case "5":
		sort.Slice(Users_array, func(i, j int) bool {
			return Users_array[i].Tel < Users_array[j].Tel
		})
	case "6":
		sort.Slice(Users_array, func(i, j int) bool {
			return Users_array[i].Desc < Users_array[j].Desc
		})
	case "7":
		break
	}
	return Users_array
}

// func SortSlice(num int) [][]string {
// 	u := Userdecode()

// 	for i := 0; i < len(u)-1; i++ {
// 		for j := 0; j < len(u)-1; j++ {
// 			if u[j][num] > slice[j+1][num] {
// 				slice[j], slice[j+1] = slice[j+1], slice[j]
// 				//fmt.Printf("%T, %s\n", slice[num], slice[num])
// 			}
// 			//fmt.Printf("%T, %s\n", slice[j][1], slice[j][1])
// 		}
// 	}
// 	//fmt.Println(slice)
// 	return slice
// }

// 定义排序函数 ，将用户按照指定的参数排序，并返回用户数组

// func User_sort() [][]string {

// 	// 输入按照哪个参数排序，并将输入的参数赋值给变量o
// 	sortKey := InputString("请输入需要排序的键值:")

// 	if len(CsvrUser()[1:]) == 0 {
// 		fmt.Println("查询内容为空")
// 		return [][]string{}
// 	} else {
// 		switch sortKey {
// 		case "1":
// 			return SortSlice(1)
// 		case "2":
// 			return SortSlice(2)
// 		case "3":
// 			return SortSlice(3)
// 		case "4":
// 			return SortSlice(4)
// 		case "5":
// 			return SortSlice(5)
// 		case "6":
// 			return SortSlice(6)
// 		default:
// 			return SortSlice(1)
// 		}

// 	}

// }

// 定义函数将排序后的用户打印

func Print_sort() {
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println(Sort_menu)

	// 将排序后的用户数组赋值给list变量
	list := User_sort()

	// 打印排序后的用户
	//fmt.Println(list)
	for _, v := range list {
		fmt.Println(v)
		fmt.Println(strings.Repeat("*", 30))
	}
}
