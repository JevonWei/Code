package main

import "fmt"

func main() {
	// yes == "Y" "y"
	var yes string
	fmt.Println("有卖西瓜的吗？(Y/N):")
	fmt.Scan(&yes)

	fmt.Println("老婆的想法")

	fmt.Println("十个包子")
	switch yes {
	case "y", "Y":
		fmt.Println("一个西瓜")
	}

	fmt.Println("老公的想法")

	switch yes {
	case "y", "Y":
		fmt.Println("一个包子")
	default:
		fmt.Println("十个包子")
	}

	// 判断成绩
	fmt.Println("------------")
	fmt.Println("判断输入的分数")

	var score int
	fmt.Println("请输入分数")
	fmt.Scan(&score)
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	case score >= 60:
		fmt.Println("D")
	default:
		fmt.Println("E")
	}


}