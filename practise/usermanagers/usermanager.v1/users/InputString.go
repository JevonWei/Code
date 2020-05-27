package users

import (
	"fmt"
	"strings"
)

// 定义用户键盘输入函数

func InputString(str string) string {
	var s string
	fmt.Print(str)
	fmt.Scan(&s)
	return strings.TrimSpace(s)
}
