package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
	"todolist/models"
	"todolist/session"
)

func TaskAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultSessionManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
	} else {
		q := r.FormValue("q")
		var tasks []models.Task
		if q == "" {
			tasks = models.GetTasks()
		} else {
			tasks = models.GetUserTasks(q)
		}

		tpl, err := template.New("tasks.html").ParseFiles("views/task/tasks.html")
		if err != nil {
			log.Println("views/task/tasks.html 渲染：", err)
		}

		tpl.Execute(w, tasks)
	}
}

func UserTaskAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultSessionManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	name := r.FormValue("name")
	log.Println("List User Task")
	tpl, err := template.New("tasks.html").ParseFiles("views/task/tasks.html")
	if err != nil {
		log.Println("views/task/tasks.html 渲染：", err)
	}
	tpl.Execute(w, models.GetUserTasks(name))

}

func TaskCreateAction(w http.ResponseWriter, r *http.Request) {
	sessionObj := session.DefaultSessionManager.SessionStart(w, r)
	if _, ok := sessionObj.Get("user"); !ok {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("create.html").ParseFiles("views/task/create.html"))
		tpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		desc := r.PostFormValue("desc")
		user := r.PostFormValue("user")

		models.CreateTask(name, user, desc)

		http.Redirect(w, r, "/task", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func TaskModifyAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err == nil {
			task, err := models.GetTaskById(id)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				tpl := template.Must(template.New("modify.html").ParseFiles("views/task/modify.html"))
				tpl.Execute(w, task)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else if r.Method == http.MethodPost {
		id, err := strconv.Atoi(r.PostFormValue("id"))

		if err != nil {
			log.Println("Task Modify Post Request error")
			w.WriteHeader(http.StatusBadRequest)
		}
		name := r.PostFormValue("name")
		desc := r.PostFormValue("desc")

		progress, err := strconv.Atoi(r.PostFormValue("progress"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		user := r.PostFormValue("user")
		status, err := strconv.Atoi(r.PostFormValue("status"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		log.Println("Modify Task Post Requrst success")

		models.ModifyTask(id, name, desc, progress, user, status)

		http.Redirect(w, r, "/task/", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func TaskDeleteAction(w http.ResponseWriter, r *http.Request) {
	if id, err := strconv.Atoi(r.FormValue("id")); err == nil {
		log.Println("发起Delete Task Request")
		models.DeleteTask(id)
	}
	http.Redirect(w, r, "/task/", http.StatusFound)
}

func init() {
	http.HandleFunc("/task/", TaskAction)
	http.HandleFunc("/task/create/", TaskCreateAction)
	http.HandleFunc("/task/modify/", TaskModifyAction)
	http.HandleFunc("/task/delete/", TaskDeleteAction)
	http.HandleFunc("/user/tasks/", UserTaskAction)

}
