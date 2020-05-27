package main

import (
	"os"
	"text/template"
)

func main() {
	tpl, _ := template.New("index.html").ParseFiles("views/index.html")
	tpl = template.Must(tpl.ParseFiles("views/danran.html"))
	tpl.Execute(os.Stdout, nil)
}
