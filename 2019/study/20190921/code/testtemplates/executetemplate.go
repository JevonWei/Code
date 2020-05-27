package main

import (
	"html/template"
	"os"
)

func main() {
	// 有block在前，define在后
	tpl := template.Must(template.ParseFiles("views/index.html", "views/jevon.html"))
	//tpl.ExecuteTemplate(os.Stdout, "index.html", nil)
	tpl.ExecuteTemplate(os.Stdout, "name", nil)

	for _, tpl := range tpl.Template() {
		fmt.Println(tpl)
	}
}
