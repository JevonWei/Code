package main

import (
	"net/http"
	_ "todolist/controllers"
	"todolist/models"
)

func main() {
	// certFile := "tools/cert.pem"
	// keyFile := "tools/key.pem"
	models.Init()
	addr := ":9999"
	http.ListenAndServe(addr, nil)
	// http.ListenAndServeTLS(addr, certFile, keyFile, nil)

}
