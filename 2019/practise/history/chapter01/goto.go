package main

import "fmt"

func main() {
	// goto 案例一
	// yes == "Y" "y"
	var yes string
	fmt.Println("有卖西瓜的吗？(Y/N):")
	fmt.Scan(&yes)

	fmt.Println("老婆的想法")
	fmt.Println("十个包子")

	if yes != "Y" && yes != "y" {
		goto END
	}

	fmt.Println("一个西瓜")

END:

	// goto 案例二
	// 100以内数字相加1...100
	result := 0
	i := 1

START:
	if i > 100 {
		goto FOREND
	}

	result += i
	i++
	goto START

FOREND:

	fmt.Println(result)


	// goto 案例三
BREAKEND:    
	for i := 0; i < 5; i++ {
		for j := 0; j < 5, j++ {
			if i * j == 9 {
				break BREAKEND     // break goto LABLE必须定义在前面
			}
		}
	}


}
