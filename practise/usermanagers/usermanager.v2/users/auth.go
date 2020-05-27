package users

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/howeyc/gopass"
)

// 系统登录认证程序

func Auth() bool {
	Passwd, err := ioutil.ReadFile(PasswordFile)
	if err == nil && len(Passwd) > 0 {
		for i := 0; i < MaxAuth; i++ {
			fmt.Printf("请输入系统密码:")
			// 密码输入
			bytes, _ := gopass.GetPasswd()

			if string(Passwd) == fmt.Sprintf("%x", md5.Sum(bytes)) {
				return true
			} else {
				if i != 2 {
					fmt.Println("密码输入错误，请重新输入")
				}
			}
		}
		fmt.Printf("密码输入%d次错误，程序退出\n", MaxAuth)
		return false
	} else {
		if os.IsNotExist(err) || len(Passwd) == 0 {
			// 文件文件不存在，初始化
			fmt.Print("请输入初始化密码:")
			bytes, _ := gopass.GetPasswd()
			ioutil.WriteFile(PasswordFile, []byte(fmt.Sprintf("%x", md5.Sum(bytes))), os.ModePerm)
			return true
		} else {
			fmt.Println("发送错误", err)
			return false
		}
	}
}
