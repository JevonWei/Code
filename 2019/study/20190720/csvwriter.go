package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Create("user.csv")

	if err == nil {
		defer file.Close()

		writer := csv.NewWriter(file)
		writer.Write([]string{"编号", "名字", "年龄"})
		writer.Write([]string{"1", "Dan", "10"})
		writer.Write([]string{"2", "Ran", "20"})
		writer.Flush()
	}
}
