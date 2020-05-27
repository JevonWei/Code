package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 随机获取100内整数
	rand.Seed(time.Now().Unix())
	var rangenumber int = rand.Intn(100)
	//fmt.Println("Jevon ",rangenumber)

	const maxGuessNum = 5
	var isOK bool

	// 用户循环输入五次
	for i := 1; i <= maxGuessNum; i++ {
		fmt.Println("---------------------- ")
		var number int
		fmt.Println("请输入一个100以内的整数: ")
		fmt.Scan(&number)
		//fmt.Printf("输入的number是: %d\n", number)

		// 判断用户输入的是否正确
		switch {
		case number == rangenumber:
			fmt.Printf("你输入的number为：%d,恭喜你，猜对了.", number)
			isOK = true
			break
		case number > rangenumber:
			fmt.Printf("你输入的number为：%d,很遗憾，你猜大了.\n你还剩%d次机会\n\n", number, maxGuessNum-i)

		case number < rangenumber:
			fmt.Printf("你输入的number为：%d,很遗憾，你猜小了.\n你还剩%d次机会\n\n", number, maxGuessNum-i)
		}
	}

	if !isOK {
		fmt.Println("太笨了，游戏结束")
	}

}
