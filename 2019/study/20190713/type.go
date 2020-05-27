package main

import "fmt"

type Counter int
type User map[string]string
type Callback func(...string)

func main() {
	var counter Counter = 20

	counter += 10
	fmt.Println(counter)

	me := make(User)
	me["name"] = "JevonWei"
	me["addr"] = "abcd"

	fmt.Println(me)
	fmt.Printf("%T, %T\n", counter, me)

	var list Callback = func(args ...string) {
		for i, v := range args {
			fmt.Println(i, v)
		}
	}

	list("a", "b", "c")

	// 定义int类型 counter2
	var counter2 int = 10
	// 将自定义Counter类型的变量counter转换为int类型
	fmt.Println(int(counter) > counter2)
	// 将int类型的变量counter2 转换为自定义Counter类型
	fmt.Println(counter > Counter(counter2))
}
