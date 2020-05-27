package utils

import (
	"math/rand"
	"time"
)

// 随机生成长度为length的字符串
func RandString(length int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWSYZ0123456789"
	count := len(letters)
	chars := make([]byte, length)
	for i := 0; i < length; i++ {
		chars[i] = letters[rand.Int()%count]
	}
	return string(chars)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
