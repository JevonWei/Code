package main

import "fmt"

func main() {
	var scores map[string]int
	fmt.Printf("%T %#v\n", scores, scores)
	fmt.Println(scores == nil)

	// 字面量
	scores = make(map[string]int)
	fmt.Printf("%T %#v\n", scores, scores)

	scores = map[string]int{"wei": 90, "zi": 56, "Jevon": 78}
	fmt.Printf("%v\n", scores)

	// 增 删 改 查
	// key
	fmt.Println(scores["wei"])
	fmt.Println(scores["dan"])

	if v, ok := scores["wei"]; ok {
		fmt.Println(v)
	}

	// 改
	scores["wei"] = 60
	fmt.Println(scores["wei"])

	// 增
	scores["dan"] = 60
	fmt.Println(scores["dan"])

	// 删除
	delete(scores, "wei")
	fmt.Println(scores)
	scores["ran"] = 100
	fmt.Println(scores["ran"])

	fmt.Println(len(scores))

	for k, v := range scores {
		fmt.Println(k, v)
	}

	// key至少可以有==, != 运算， bool类型，整数， float，字符串， 数组
	// value可以是任意类型
	users := map[string]map[string]string{"dan": {"quyu": "hn", "phone": "123"}}
	//users = map[string]map[string]string{"dan":{"quyu":"hn", "phone":"123"}}
	fmt.Printf("%T %#v\n", users, users)
	fmt.Println(users["dan"]["quyu"])
	if _, ok := users["ran"]; !ok {
		users = map[string]map[string]string{"ran": {"quyu": "sh", "phone": "0987"}}
	}
	fmt.Printf("%T %#v\n", users, users)

	delete( users["ran"], "phone")
	fmt.Printf("%T %#v\n", users, users)
}

/*
func main() {
	var scores map[string]int

	fmt.Printf("%T %#v\n", scores, scores)
	fmt.Println(scores == nil)

	// 字面量

	scores = make(map[string]int)
	fmt.Println(scores)

	// scores = map[string]int{}
	scores = map[string]int{"Jevon": 8, "dan": 3, "Wei": 9}
	fmt.Println(scores)


	// 增，删，改，查
	// key
	fmt.Println(scores["Wei"])
	fmt.Println(scores["dan"])
	//_, ok := scores["dan"]
	v, ok := scores["ran"]
	if ok {
		fmt.Println(v)
	}


	if v, ok := scores["dan"]; ok {
		fmt.Println(v)
	}

	scores["dan"] = 7
	fmt.Println(scores)
	scores["ran"] = 7
	fmt.Println(scores)

	// 删除键值
	delete(scores, "dan")
	fmt.Println(scores)

	fmt.Println(len(scores))
	for k, v := range scores {
		fmt.Println(k, v)
	}

	// key至少有==, != 运算, bool, 整数， 字符串， 数组
	// value => 为任意类型 slice map

	var users map[string]map[string]string
	users = map[string]map[string]string{"dan":{"1":"a","2":"b"}}
	users["wei"] = map[string]string{"4":"c","5":"d"}
	fmt.Printf("%T, %#v\n", users, users)

	fmt.Println(users["dan"]["1"])
}
*/
