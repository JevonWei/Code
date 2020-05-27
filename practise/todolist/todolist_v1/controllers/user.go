package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"todolist/models"
	"todolist/session"
)

func Menu(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login/", http.StatusFound)
}

func LoginAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl, err := template.New("login.html").ParseFiles("views/user/login.html")
		if err != nil {
			log.Panic("Login Auth Request Error")
		}
		// tpl.Execute(w, nil)
		tpl.Execute(w, struct {
			Name  string
			Error error
		}{})

	} else if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		password := models.PasswdMd5(r.PostFormValue("password"))
		if r.PostFormValue("login") != "" {

			if ok, err := models.LoginAuth(name, password); ok {
				user, _ := models.GetUserByName(name)
				sessionObj := session.DefaultSessionManager.SessionStart(w, r)
				sessionObj.Set("user", user)

				if models.SuperUser(name) {
					http.Redirect(w, r, "/userlist/", http.StatusFound)
				} else {
					http.Redirect(w, r, "/task/", http.StatusFound)
				}
			} else {
				tpl, _ := template.New("login.html").ParseFiles("views/user/login.html")

				tpl.Execute(w, struct {
					Name  string
					Error error
				}{name, err})
				// fmt.Fprintf(w, "%s登录验证失败\n", name)
				http.Redirect(w, r, "/login/", http.StatusFound)
			}
		} else {
			http.Redirect(w, r, "/register/", http.StatusFound)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl, err := template.New("register.html").ParseFiles("views/user/register.html")
		if err != nil {
			log.Panic("Register Request Error")
		}
		tpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		name := r.PostFormValue("name")

		password := models.PasswdMd5(r.PostFormValue("password"))
		desc := r.PostFormValue("desc")
		tel := r.PostFormValue("tel")
		addr := r.PostFormValue("addr")

		if models.CheckUserName(name) {
			super := false
			models.Register(name, password, desc, tel, addr, super)
			http.Redirect(w, r, "/login/", http.StatusFound)
		} else {
			fmt.Fprintf(w, "用户已存在，请重新注册")
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func UserlistAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultSessionManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
	} else {
		q := strings.TrimSpace(r.FormValue("q"))
		users := models.GetSearchUser(q)
		tpl := template.Must(template.New("list.html").ParseFiles("views/user/list.html"))

		context := struct {
			Q     string
			Users []models.User
		}{q, users}

		tpl.Execute(w, context)
	}

}

func UserCreateAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("create.html").ParseFiles("views/user/create.html"))
		tpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		password := r.PostFormValue("password")
		tel := r.PostFormValue("tel")
		addr := r.PostFormValue("addr")
		super := r.PostFormValue("super")
		desc := r.PostFormValue("desc")

		superbool, err := strconv.ParseBool(super)
		if err != nil {
			log.Panic("User Create Super值获取失败")
		}

		if models.CheckUserName(name) {
			models.CreateUser(name, password, tel, addr, desc, superbool)
			http.Redirect(w, r, "/userlist/", http.StatusFound)
		} else {
			fmt.Fprintf(w, "用户已存在，请重新注册")
		}

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func LogoutAction(w http.ResponseWriter, r *http.Request) {
	session.DefaultSessionManager.SessionDestory(w, r)
	http.Redirect(w, r, "/login/", http.StatusFound)
}

func init() {
	http.HandleFunc("/", Menu)
	http.HandleFunc("/login/", LoginAuth)
	http.HandleFunc("/logout/", LogoutAction)
	http.HandleFunc("/register/", Register)
	http.HandleFunc("/userlist/", UserlistAction)
	http.HandleFunc("/user/create/", UserCreateAction)
}
