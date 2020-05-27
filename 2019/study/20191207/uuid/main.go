package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	// 生成uuid
	fmt.Println(uuid.New().String())
}
