package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	urllib "net/url"
	"os"
)

func main() {
	url := "http://localhost:8080/request/"

	response, err := http.Get(url)
	if err == nil {
		fmt.Println(response.Proto, response.Status)
		fmt.Println(response.Header)

		io.Copy(os.Stdout, response.Body)
	}

	json := bytes.NewReader([]byte(`{"name" : "JevonWei", "password" : "121212"}`))
	response, err = http.Post(url, "application/json", json)
	if err == nil {
		fmt.Println(response.Proto, response.Status)
		fmt.Println(response.Header)

		io.Copy(os.Stdout, response.Body)
	}

	params := make(urllib.Values)
	params.Add("name", "Jevon")
	params.Add("password", "123456")
	response, err = http.PostForm(url, params)
	if err == nil {
		fmt.Println(response.Proto, response.Status)
		fmt.Println(response.Header)

		io.Copy(os.Stdout, response.Body)
	}

}
