package utils

import (
	"html/template"
	"io"
)

func Render(w io.Writer, name string, files []string, context interface{}) {
	tpl := template.Must(template.New(name).ParseFiles(files...))
	tpl.Execute(w, context)
}
