package main

import (
	"html/template"
	"os"
)

func main() {
	tpl := template.Must(template.New("index.html").ParseFiles("views/index.html"))
	tpl = template.Must(tpl.ParseFiles("views/jevon.html"))
	tpl.Execute(os.Stdout, nil)
}
