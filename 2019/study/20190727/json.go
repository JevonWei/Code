package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

func main() {
	/*
		json.Marshal   序列化
		json.Unmarshal 反序列化
	*/

	names := []string{"Dan", "Ran", "Jevon"}
	users := []map[string]string{{"name": "Dan", "Addr": "shanghai"}, {"name": "Ran", "Addr": "Henan"}}

	bytes, err := json.MarshalIndent(names, "", "\t")
	//bytes, err := json.Marshal(names)
	if err == nil {
		//fmt.Println(bytes)
		fmt.Println(string(bytes))
	}

	var names02 []string
	err = json.Unmarshal(bytes, &names02)
	fmt.Println(err)
	fmt.Println(names02)

	bytes, err = xml.MarshalIndent(&users, " ", "    ")
	//bytes, err = json.Marshal(users)
	//bytes, err = xml.Marshal(users)
	if err == nil {
		//fmt.Println(bytes)
		fmt.Println(string(bytes))
	}

	var users02 []map[string]string
	err = xml.Unmarshal(bytes, &users02)
	fmt.Println(err)
	fmt.Println(users02)

	// 判断json的格式是否正确
	fmt.Println(json.Valid([]byte("[]")))
	//fmt.Println(json.Valid([]byte("[]x")))
}
