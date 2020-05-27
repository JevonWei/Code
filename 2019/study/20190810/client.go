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
	url := "http://localhost:9999/request/"

	// head
	response, err := http.Get(url)

	if err == nil {
		fmt.Println(response.Proto, response.Status)
		fmt.Println(response.Header)
		io.Copy(os.Stdout, response.Body)
	}

	json := bytes.NewReader([]byte(`{"name":"Jevon", "passwd":"12345"}`))
	response, err = http.Post(url, "application/json", json)
	if err == nil {
		fmt.Println(response.Proto, response.Status)
		fmt.Println(response.Header)
		io.Copy(os.Stdout, response.Body)
	}

	params := make(urllib.Values)
	params.Add("name", "Jevon")
	params.Add("password", "123321")

	response, err = http.PostForm(url, params)
	if err == nil {
		fmt.Println(response.Proto, response.Status)
		fmt.Println(response.Header)
		io.Copy(os.Stdout, response.Body)
		//response.Write(os.Stdout)
	}

}
