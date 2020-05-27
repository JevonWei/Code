package models

import (
	"database/sql"
	"time"
)

type Task struct {
	Id           int        `json:"id"`
	Name         string     `json:"name"`
	Progress     int        `json:"progress"`
	User         string     `json:"user"`
	Desc         string     `json:"desc"`
	Status       int        `json:"staus"`
	CreateTime   *time.Time `json:"create_time"`
	CompleteTime *time.Time `json:"complete_time"`
}

func GetTasks() []Task {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select id, name, progress, user, `desc`, status, create_time, complete_time from tasks")
	if err != nil {
		panic(err)
	}

	tasks := make([]Task, 0)
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.Id, &task.Name, &task.Progress, &task.User, &task.Desc, &task.Status, &task.CreateTime, &task.CompleteTime); err == nil {
			tasks = append(tasks, task)
		}

	}
	return tasks
}

func CreateTask(name, user, desc string) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	defer db.Close()
	_, err = db.Exec("insert into tasks(name, user, `desc`, create_time) values(?, ?, ?, ?)", name, user, desc, time.Now())

	if err != nil {
		panic(err)
	}
}

func GetTaskById(id int) (Task, error) {
	var task Task
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return task, nil
	}

	if err := db.Ping(); err != nil {
		return task, nil
	}

	defer db.Close()

	row := db.QueryRow("select id, name, progress, user, `desc`, status, create_time, complete_time from tasks where id=?", id)
	err = row.Scan(&task.Id, &task.Name, &task.Progress, &task.User, &task.Desc, &task.Status, &task.CreateTime, &task.CompleteTime)

	return task, err
}

func ModifyTask(id int, name, desc string, progress int, user string, status int) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	defer db.Close()

	_, err = db.Exec("update tasks set name=?, `desc`=?, progress=?, user=?, status=? where id = ?", name, desc, progress, user, status, id)

	if err != nil {
		panic(err)
	}
}

func DeleteTask(id int) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	defer db.Close()

	_, err = db.Exec("delete from tasks where id = ?", id)

	if err != nil {
		panic(err)
	}
}
