package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"todolist/models"
)

func TaskAction(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("task.html").ParseFiles("views/task.html"))
	tpl.Execute(w, models.GetTasks())
}

func TaskCreateAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl := template.Must(template.New("create_tast.html").ParseFiles("views/create_task.html"))
		tpl.Execute(w, nil)

	} else if r.Method == http.MethodPost {
		name := r.PostFormValue("name")
		desc := r.PostFormValue("desc")
		user := r.PostFormValue("user")

		models.CreateTask(name, user, desc)
		http.Redirect(w, r, "/", http.StatusFound)
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
				w.WriteHeader(http.StatusBadGateway)
			} else {
				tpl := template.Must(template.New("modify_task.html").ParseFiles("views/modify_task.html"))
				tpl.Execute(w, task)
			}
		} else {
			w.WriteHeader(http.StatusBadGateway)
		}

	} else if r.Method == http.MethodPost {
		id, err := strconv.Atoi(r.PostFormValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		name := r.PostFormValue("name")
		desc := r.PostFormValue("desc")
		progress, err := strconv.Atoi(r.PostFormValue("progress"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		user := r.PostFormValue("user")
		status := r.PostFormValue("status")

		models.ModifyTask(id, name, desc, progress, user, status)

		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func TaskDeleteAction(w http.ResponseWriter, r *http.Request) {
	if id, err := strconv.Atoi(r.FormValue("id")); err == nil {
		models.DeleteTask(id)
	} else {
		fmt.Println(err)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func init() {
	http.HandleFunc("/", TaskAction)
	http.HandleFunc("/task/create/", TaskCreateAction)
	http.HandleFunc("/task/modify/", TaskModifyAction)
	http.HandleFunc("/task/delete/", TaskDeleteAction)
}
