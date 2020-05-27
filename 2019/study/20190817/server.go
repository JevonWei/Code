package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	addr := ":9999"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		txt := `
		<!DOCTYPE html>
		<html>
			<head>
				<meta charset="utf-8" />
				<title> JevonWei</title>
			</head>
			<body>
				<form action="/register/?a=b&c=d" method="POST" enctype="multipart/form-data">
					<label> 用户名:</label><input name="username" type="text" />
					<label>密码:</label><input name="password" type="password" />
					<label>头像:</label><input name="img" type="file" />
					<input type="submit" value="注册" />  
				</form>      
			</body>
		</html>
		`
		fmt.Fprint(w, txt)
	})

	http.HandleFunc("/register/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Method, request.URL)

		// 对URL参数解析
		// request.ParseForm()
		request.ParseMultipartForm(1024) // multipart/form-data时解析值，需设置解析最大内存为1024字节

		fmt.Println(request.Form)                // 获取url和body所有数据
		fmt.Println(request.PostForm)            // 只获取body数据
		fmt.Println(request.MultipartForm)       // 获取body中的所有数据，包括文件类型的数据
		fmt.Println(request.MultipartForm.Value) // 获取body中的value数据
		fmt.Println(request.MultipartForm.File)  // 获取body中的文件类型的数据

		fmt.Printf("%T\n", request.MultipartForm.File["img"][0])      // 打印文件的类型
		if file, header, err := request.FormFile("img"); err == nil { // request.FormFile("img")获取上传的文件
			fmt.Printf("%T, %T, %v\n", file, header, err)
			fmt.Println(header.Filename, header.Header, header.Size)
			fmt.Println(header.Header["Content-Type"]) // 获取上传文件的类型
			newFile, err := os.Create("1.txt")
			if err == nil {
				defer newFile.Close()
				io.Copy(newFile, file)
			}
			fmt.Println("ok")
		} else {
			fmt.Println(err)
		}

		// fmt.Println(request.FormValue("username"))     // 自动调用parseForm
		// fmt.Println(request.PostFormValue("username")) // 自动调用parseForm
		// fmt.Println(request.Form["username"])
		// fmt.Println(request.Form["password"])
		// fmt.Println(request.Form["username"][0]
		// fmt.Println(request.Form["password"][0])
		// fmt.Println(request.Form.Get("username"))
		// fmt.Println(request.Form.Get("password"))
		//fmt.Fprintf(response, "ok")
		http.Redirect(response, request, "/login/", 301)
	})

	http.HandleFunc("/login/", func(respnse http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(respnse, "用户登录")
	})

	http.ListenAndServe(addr, nil)
}
