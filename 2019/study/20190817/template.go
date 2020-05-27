package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cxt := `
		显示数据:{{ . }}
		显示数据:{{ index . 1 }}
		` 

		tpl, err := template.New("tpl").Parse(cxt)
		if err == nil {
			tpl.Execute(w, []string{"Hello World", "Jevon"})
		} else {
			fmt.Println(err)
		}
	})

	http.HandleFunc("v1", func(w http.ResponseWriter, r *http.Request) {
		cxt := `
		显示数据:{{ . }}
		显示数据:{{ .A }}
		`

		tpl, err := template.New("tpl").Parse(cxt)
		if err == nil {
			tpl.Execute(w, map[string]string{"A": "Hello World", "B": "Jevon"})
		} else {
			fmt.Println(err)
		}
	})

	http.HandleFunc("/v2", func(w http.ResponseWriter, r *http.Request) {
		cxt := `
		显示数据:{{ . }}
		显示数据:{{ .A }}
		{{ if .}}
		True
		{{ else }}
		False
		{{ end }}	
		
		`

		tpl, err := template.New("tpl").Parse(cxt)
		if err == nil {
			tpl.Execute(w, true)
		} else {
			fmt.Println(err)
		}
	})

	http.HandleFunc("/v3", func(w http.ResponseWriter, r *http.Request) {
		cxt := `
		显示数据:{{ . }}
		{{ if ge . 90}}
		优秀
		{{ else if ge . 60}}
		良好
		{{ else }}
		差	
		{{ end }}
		`

		// ge >=
		// gt >
		// le <-
		// lt <
		// eq ==
		// ne !

		tpl, err := template.New("tpl").Parse(cxt)
		if err == nil {
			tpl.Execute(w, 80)
		} else {
			fmt.Println(err)
		}
	})

	http.HandleFunc("/v4", func(w http.ResponseWriter, r *http.Request) {
		cxt := `
		显示数据:{{ . }}

		{{ range . }}
		元素：{{ . }}
		{{ end }}
		`

		tpl, err := template.New("tpl").Parse(cxt)
		if err == nil {
			tpl.Execute(w, []string{"Jevon", "Dan", "Ran"})
		} else {
			fmt.Println(err)
		}
	})

	// 遍历map，定义变量
	http.HandleFunc("/v5", func(w http.ResponseWriter, r *http.Request) {
		cxt := `
		{{ $name := "JevonWei" }}
		显示数据:{{ . }} {{ $name }}
		{{ range $key, $value := . }}
		元素：{{ . }} {{ $key }} {{ $value }}
		{{ end }}
		`

		tpl, err := template.New("tpl").Parse(cxt)
		if err == nil {
			tpl.Execute(w, map[string]string{"A": "Hello World", "B": "Jevon"})
		} else {
			fmt.Println(err)
		}
	})

	// 自定义函数
	http.HandleFunc("/v6", func(w http.ResponseWriter, r *http.Request) {
		cxt := `
		{{ index . 0}}
		{{ 0|index . }}
		{{ len . }}
		{{ .|len }}
		{{ printf "%T" . }}
		{{ .|printf "%T" }}

		自定义函数： {{ "abc"|upper }}
		{{ "abc"|upper2}}
		{{ upper "abc" }}
		{{ upper2 "abc" }}
		`

		// 自定义函数
		funcs := template.FuncMap{
			"upper": func(txt string) string {
				return strings.ToUpper(txt)
			},
			"upper2": strings.ToUpper,
		}

		tpl, err := template.New("tpl").Funcs(funcs).Parse(cxt)
		if err == nil {
			tpl.Execute(w, []string{"Hello World", "Jevon"})
		} else {
			fmt.Println(err)
		}
	})

	http.HandleFunc("/v7", func(w http.ResponseWriter, r *http.Request) {

		tpl, err := template.New("tpl.html").ParseFiles("tpl.html")
		if err == nil {
			tpl.Execute(w, "map[string]map[string]string{"tasks": {"a": "AA, "b": "BB", "c": "CC"}})
		} else {
			fmt.Println(err)
		}
	})

	http.ListenAndServe(":9999", nil)
}
