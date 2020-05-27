package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Compare("abc", "bac"))                              // 比较字符串
	fmt.Println(strings.Contains("abc", "ab"))                              // 字符串是否包含
	fmt.Println(strings.ContainsAny("abc", "ae"))                           // 包含字符串中的任意一个字符
	fmt.Println(strings.Count("abcabcad", "ab"))                            // 获取字符出现的次数
	fmt.Println("%q\n", strings.Fields("abc def\neeee\raaaa\taa\n\asdfxx")) // 将字符串安装空白字符分割(空格，\n, \r, \t, \f, \v)

	fmt.Println(strings.HasPrefix("abcabcad", "ab"))   // 判断是否以字符串开头
	fmt.Println(strings.HasSuffix("abcabcad", "csda")) // 判断是否以字符串结尾
	fmt.Println(strings.Index("abcabcad", "bc"))       // 找出字符串出现的位置
	fmt.Println(strings.Index("abcabcad", "cad"))
	fmt.Println(strings.LastIndex("abcabcad", "cad")) // 获取字符串最后一个出现的位置

	fmt.Println(strings.Split("abcabcad;asd;asf", ";"))           // 将字符串按特定符号分割
	fmt.Println(strings.Join([]string{"abc", "def", "eee"}, ":")) //将字符串按符号连接在一起

	fmt.Println(strings.Repeat("abc", 3))                        // 将字符串重复打印n次
	fmt.Println(strings.Replace("abcabcabcac", "ab", "xxx", 1))  // 将一个ab替换为xxx
	fmt.Println(strings.Replace("abcabcabcac", "ab", "xxx", -1)) // 将所有的ab替换为xxx
	fmt.Println(strings.ReplaceAll("abcabcabcac", "ab", "xxx"))  // 将所有的ab替换为xxx

	fmt.Println(strings.ToLower("abcdefg"))  // 将所有字符串转换为小写
	fmt.Println(strings.ToUpper("ABCDdefg")) // 将所有字符串转换为大写
	fmt.Println(strings.Title("jevonwei"))   // 将字符串的首字母大写

	fmt.Println(strings.Trim("abcdefgzxaszxxz", "zx"))    // 将特定字符zx去掉
	fmt.Println(strings.TrimSpace(" abcdefg  XXX \n \r")) // 将空白字符去掉
}
