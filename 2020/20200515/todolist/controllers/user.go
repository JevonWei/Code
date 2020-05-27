package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"todolist/models"
	"todolist/session"
	"todolist/utils"
)

func UserAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/user/login/", http.StatusFound)
		return
	}

	q := strings.TrimSpace(r.FormValue("q"))

	// tpl := template.Must(template.New("user.html").ParseFiles("views/user/user.html"))

	// tpl.Execute(w, struct {
	// 	Query string
	// 	Users []models.User_str
	// }{q, models.GetUsers(q)})

	context := struct {
		Query string
		Users []models.User_str
	}{q, models.GetUsers(q)}
	utils.Render(w, "base.html", []string{"views/layouts/base.html", "views/user/user.html"}, context)
}

func LoginAction(w http.ResponseWriter, r *http.Request) {
	content := struct {
		Name  string
		Error string
	}{"", ""}

	if r.Method == http.MethodGet {
		// tpl := template.Must(template.New("login.html").ParseFiles("views/user/login.html"))
		// tpl.Execute(w, nil)
		utils.Render(w, "login.html", []string{"views/user/login.html"}, content)
	} else if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		password := r.PostFormValue("password")

		user, err := models.GetUserByName(name)
		if err != nil || !user.ValidatePassword(password) {
			// tpl := template.Must(template.New("login.html").ParseFiles("views/user/login.html"))
			// tpl.Execute(w, struct {
			// 	Name  string
			// 	Error string
			// }{name, "用户名或密码错误"})
			content.Name = name
			content.Error = "用户名或密码错误"
			utils.Render(w, "login.html", []string{"views/user/login.html"}, content)
		} else {
			sessionObj := session.DefaultManager.SessionStart(w, r)
			//登陆成功
			sessionObj.Set("user", user)
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
}

func LogoutAction(w http.ResponseWriter, r *http.Request) {
	session.DefaultManager.SessionDestory(w, r)
	http.Redirect(w, r, "/user/login/", http.StatusFound)
}

func UserCreateAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/user/login/", http.StatusFound)
		return
	}

	var context interface{}
	// var context struct {
	// 	Errors   models.Errors
	// 	Name     string
	// 	Password string
	// 	Birthday string
	// 	Tel      string
	// 	Addr     string
	// 	Desc     string
	// }

	if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		password := r.PostFormValue("password")
		birthday := r.PostFormValue("birthday")
		tel := r.PostFormValue("tel")
		desc := r.PostFormValue("desc")
		addr := r.PostFormValue("addr")
		sex, _ := strconv.ParseBool(r.PostFormValue("sex"))

		// errors := models.ValidateCreateUser(name, password, birthday, tel, addr, desc)
		//if len(errors) == 0 {
		// var err models.Errors
		// if errors == err {
		// 	models.CreateUser(name, password, birthday, tel, addr, desc, sex)
		// 	http.Redirect(w, r, "/users/", http.StatusFound)
		// 	return
		// }

		errors := models.ValidateCreateUser(name, password, birthday, tel, addr, desc)
		if len(errors) == 0 {
			models.CreateUser(name, password, birthday, tel, addr, desc, sex)
			http.Redirect(w, r, "/users/", http.StatusFound)
			return
		}

		context = struct {
			Errors   map[string]string
			Name     string
			Password string
			Birthday string
			Tel      string
			Addr     string
			Desc     string
		}{errors, name, password, birthday, tel, addr, desc}
	}

	// tpl := template.Must(template.New("create.html").ParseFiles("views/user/create.html"))
	// tpl.Execute(w, context)
	utils.Render(w, "base.html", []string{"views/layouts/base.html", "views/user/create.html"}, context)
}

func init() {
	http.HandleFunc("/users/", UserAction)
	http.HandleFunc("/users/create/", UserCreateAction)
	http.HandleFunc("/user/login/", LoginAction)
	http.HandleFunc("/user/logout/", LogoutAction)
}
