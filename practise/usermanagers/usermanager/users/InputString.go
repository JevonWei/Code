package users

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 定义用户键盘输入函数

func InputString(str string) string {
	// var s string
	fmt.Print(str)
	// fmt.Scan(&s)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Text()
	return strings.TrimSpace(scanner.Text())
}
