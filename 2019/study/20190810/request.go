package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	//url := "http://localhost:9999/request/"
	//url := "https://www.baidu.com"
	url := "https://202.89.233.100/"

	// 创建request对象
	request, _ := http.NewRequest("DELETE", url, nil)

	// 跳过https的ssl认证
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// 创建Client对象
	client := &http.Client{Transport: transport}
	// 将request请求对象传入client.Do()函数,返回*response类型
	response, err := client.Do(request)

	if err == nil {
		fmt.Println(response.Proto, response.Status)
		fmt.Println(response.Header)
		io.Copy(os.Stdout, response.Body)
		//response.Write(os.Stdout)
	} else {
		fmt.Println(err)
	}

}
