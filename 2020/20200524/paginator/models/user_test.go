package models

import (
	"fmt"
	"testing"
)

func TestGenerateRandomTime(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(generateRandomTime().Format("2006-01-02 15:04:05"))
	}
}

func TestGenerateRandomData(t *testing.T) {
	fmt.Println(generateRandomData(10))
}
