package main

import (
	"net/http"
	_ "todolist/controllers"
	"todolist/models"
)

func main() {
	models.Init()
	addr := ":9999"
	http.ListenAndServe(addr, nil)

}
