package controllers

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"todolist/models"
)

func User(w http.ResponseWriter, r *http.Request) {
	fmt.Println("JevonWei")
	http.Redirect(w, r, "/login", http.StatusFound)
}

func ModifyPasswd(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		tpl := template.Must(template.New("modify_passwd.html").ParseFiles("views/user/modify_passwd.html"))
		tpl.Execute(w, nil)
	}
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("login.html").ParseFiles("views/user/login.html"))
		tpl.Execute(w, nil)
		//http.Redirect(w, r, "/", http.StatusFound)
	} else if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		passwd := r.PostFormValue("passwd")

		passwd_md5 := fmt.Sprintf("%x", md5.Sum([]byte(passwd)))

		if r.PostFormValue("Login") != "" {

			if Verify, err := models.AccountVerify(name, passwd_md5); Verify {

				http.Redirect(w, r, "/user", http.StatusFound)
			} else {

				http.Redirect(w, r, "/login/", http.StatusBadRequest)
				w.Write([]byte("用户验证失败，请重新登录"))
				fmt.Println(err)
				// fmt.Printf("%T, %v\n", errors.New("Account not Exist"), errors.New("Account not Exist"))
				// fmt.Printf("%T, %v\n", err, err)
				// if err == errors.New("Verify Failed") {
				// 	fmt.Printf("%T, %v\n", err, err)
				// }
				//w.Write([]byte(string(err)))
			}
		} else if r.PostFormValue("Register") != "" {
			err := models.AccountCreate(name, passwd_md5)
			if err == nil {
				w.Write([]byte("账号注册成功，请使用该账号重新登录"))
				http.Redirect(w, r, "/login", http.StatusFound)
			} else {
				w.Write([]byte("Account Exist"))
				w.Write([]byte("\n"))
				w.Write([]byte("账号已存在，请重新注册"))
				http.Redirect(w, r, "/login", http.StatusBadRequest)
				//w.Write([]byte(err))
				// fmt.Printf("%T, %v\n", errors.New("Account Exist"), errors.New("Account Exist"))
				// fmt.Printf("%T, %v\n", err, err)
				// if err == errors.New("Account Exist") {
				// 	fmt.Printf("%T, %v\n", err, err)
				// }

			}

		}

		// models.AccountCreate(name, passwd)
		// models.GetAccount()

	}

	//tpl.Execute(w, []User{})
}

func UserAction(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("user.html").ParseFiles("views/user/user.html"))
	tpl.Execute(w, models.GetUsers())
	//tpl.Execute(w, []User{})
}

func UserCreateAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("create_user.html").ParseFiles("views/user/create_user.html"))
		tpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		addr := r.PostFormValue("addr")
		tel := r.PostFormValue("tel")
		birthday := r.PostFormValue("birthday")
		desc := r.PostFormValue("desc")

		models.CreateUser(name, birthday, addr, tel, desc)

		http.Redirect(w, r, "/user", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func UserModifyAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err == nil {
			user, err := models.GetUserById(id)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				//fmt.Println(user.Birthday)
				tpl := template.Must(template.New("modify_user.html").ParseFiles("views/user/modify_user.html"))
				tpl.Execute(w, user)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}

	} else if r.Method == http.MethodPost {
		id, err := strconv.Atoi(r.PostFormValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		name := r.PostFormValue("name")
		desc := r.PostFormValue("desc")
		addr := r.PostFormValue("addr")
		// progress, err := strconv.Atoi(r.PostFormValue("progress"))
		// if err != nil {
		// 	w.WriteHeader(http.StatusBadRequest)
		// }
		tel := r.PostFormValue("tel")
		birthday := r.PostFormValue("birthday")

		models.ModifyUser(id, name, birthday, addr, tel, desc)

		http.Redirect(w, r, "/user", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func UserDeleteAction(w http.ResponseWriter, r *http.Request) {
	if id, err := strconv.Atoi(r.FormValue("id")); err == nil {
		models.DeleteUser(id)
	} else {
		fmt.Println(err)
	}
	http.Redirect(w, r, "/user", http.StatusFound)
}

func init() {
	http.HandleFunc("/", User)
	http.HandleFunc("/login", UserLogin)
	http.HandleFunc("/passwd/modify/", ModifyPasswd)

	http.HandleFunc("/user", UserAction)
	http.HandleFunc("/user/create/", UserCreateAction)
	http.HandleFunc("/user/modify/", UserModifyAction)
	http.HandleFunc("/user/delete/", UserDeleteAction)
}
