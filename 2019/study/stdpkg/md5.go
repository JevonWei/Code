package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	bytes := md5.Sum([]byte("i am jevon"))

	// 16进制显示md5加密
	fmt.Printf("字符串显示1：%x\n", md5.Sum([]byte("i an Jevon")))

	// 转换为16进制字符串
	x := fmt.Sprintf("%X\n", md5.Sum([]byte("i an Jevon")))
	fmt.Printf("字符串显示2:%x\n", x)

	// 16进制字符串显示md5
	fmt.Printf("字符串显示3:%X\n", hex.EncodeToString(bytes[:]))

	m := md5.New()
	m.Write([]byte("i am"))
	m.Write([]byte("Jevon"))
	fmt.Printf("字符串显示4:%X\n", m.Sum(nil)) // 最后不再向m中传入数据，故使用m.Sum(nil)

}
