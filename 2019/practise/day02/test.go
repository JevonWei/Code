package main

import "fmt"

func main() {
	users := []map[string]string{{"name": "Dan", "Addr": "shanghai"}, {"name": "Ran", "Addr": "Henan"}}
	for _, user := range users {
		fmt.Printf("%T, %v\n", user, user)
	}
	fmt.Printf("%T\n", users)
}
