package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World, %s!", r.URL.Path[1:])
}

type Test03 struct{}

func (t Test03) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Time: %s", time.Now().Format("2006-01-02 15:04:05"))
}

func main() {
	http.HandleFunc("/aaa", handler)
	// http.HandleFunc("/test01", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hi Jevon"))
	// })

	http.Handle("/test03/", Test03{})
	http.Handle("/test04/", &Test03{})

	http.HandleFunc("/request/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.UserAgent())
		fmt.Println(r.Referer())
		fmt.Println(r.Method, r.URL, r.Proto)
		fmt.Println(r.Header)

		fmt.Println("响应体:")

		// bytes := make([]byte, 1024)
		// n, _ := r.Body.Read(bytes)
		// fmt.Println(string(bytes[:n]))

		io.Copy(os.Stdout, r.Body)
		// w.Header().Add("Host", "Jevon")
		w.Write([]byte("Request"))

	})

	http.Handle("/", http.FileServer(http.Dir(".")))

	http.ListenAndServe(":8080", nil)
}
