package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 随机获取100内整数
	rand.Seed(time.Now().Unix())
	rand.Seed(time.Now().Unix())
	const maxGuessNum = 5
	var isOK bool

	for {
		var rangenumber int = rand.Intn(100)
		fmt.Printf("随机数为: %d", rangenumber)
		for i := 1; i <= maxGuessNum; i++ {
			fmt.Println("---------------------- ")
			var number int
			fmt.Println("请输入一个100以内的整数: ")
			fmt.Scan(&number)

			switch {
			case number == rangenumber:
				fmt.Printf("你输入的number为：%d,恭喜你，猜对了.\n", number)
				isOK = true
				i = maxGuessNum
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
		var text string
		fmt.Print("重新开始游戏吗？(y/n)")
		fmt.Scan(&text)
		if text != "y" {
			break
		}
	}
}
