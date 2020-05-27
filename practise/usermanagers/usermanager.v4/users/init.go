package users

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

// 初始化程序，定义常量及打印Menu信息

// 定义密码输入次数及系统密码的md5值
const MaxAuth = 3

// const Passwd = "bbbca3930644623c279a38c3a22c4cb5"
const (
	PasswordFile = ".password"
	UserFile     = "users.xml"
)

// 定义用户的结构体
type User struct {
	ID       int       `xml:"id,attr"`
	Name     string    `xml:"name"`
	Birthday time.Time `xml:"birthday"`
	Addr     string    `xml:"addr"`
	Tel      string    `xml:"tel"`
	Desc     string    `xml:"desc"`
}

type UserMap map[int]User
type xmlMapEntry struct {
	XMLName xml.Name `xml:"root"`
	Users   User     `xml:"user"`
}

// 定义用户map
// var Users map[int]User = map[int]User{}

// 定义系统Menu信息
var Info string = `
1. 显示
2. 查询
3. 添加
4. 修改
5. 删除
6，修改密码
7. 退出
`

var Sort_Menu string = `
1. ID
2. Name
3. Birthday
4. Addr
5. Tel
6. Desc
`

func Init() {
	fmt.Println("欢迎登录用户管理系统")

}

// 定义用户Title信息
func Print_Title() {
	title := fmt.Sprintf("%-5s|%-10s|%-15s|%-10s|%-15s|%-15s", "ID", "Name", "Birthday", "Tel", "Addr", "Desc")
	fmt.Println(title)
	fmt.Println((strings.Repeat("-", len(title))))
}
