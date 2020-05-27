package main

import (
	"html/template"
	"os"
)

func main() {
	// 有block在前，define在后
	tpl := template.Must(template.ParseFiles("views/include.html", "views/jevon.html"))
	tpl.ExecuteTemplate(os.Stdout, "include.html", nil)
}
