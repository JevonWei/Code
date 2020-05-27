package main

import (
	"os"
	"text/template"
)

func main() {
	tpl, _ := template.New("index.html").ParseFiles("views/index.html", "views/danran.html")
	tpl.Execute(os.Stdout, nil)
}
