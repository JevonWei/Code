package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/set/", func(response http.ResponseWriter, request *http.Request) {
		// tpl := template.Must(template.New("login.html").ParseFiles("login.html"))
		// tpl.Execute(response, nil)

		cookie := http.Cookie{
			Name:     "tag",
			Value:    "1",
			Path:     "/",
			MaxAge:   60,
			Expires:  time.Now().Add(time.Second * 60),
			HttpOnly: true,
		}
		http.SetCookie(response, &cookie)

	})

	http.HandleFunc("/get/", func(response http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie("tag")
		fmt.Println(cookie, err)
		response.Write([]byte("ok"))
	})

	addr := "0.0.0.0:9998"
	http.ListenAndServe(addr, nil)
}
