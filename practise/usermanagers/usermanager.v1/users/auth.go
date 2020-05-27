package users

import (
	"crypto/md5"
	"fmt"

	"github.com/howeyc/gopass"
)

// 系统登录认证程序

func Auth() bool {
	for i := 0; i < MaxAuth; i++ {
		fmt.Printf("请输入系统密码:")
		// 密码输入
		bytes, _ := gopass.GetPasswd()

		if Passwd == fmt.Sprintf("%x", md5.Sum(bytes)) {
			return true
		} else {
			if i != 2 {
				fmt.Println("密码输入错误，请重新输入")
			}
		}
	}
	fmt.Printf("密码输入%d次错误，程序退出\n", MaxAuth)
	return false
}
