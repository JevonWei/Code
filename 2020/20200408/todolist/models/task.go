package models

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

type Task struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Progress int    `json:"progress"`
	User     string `json:"user"`
	Desc     string `json:"desc"`
	Status   string `json:"status"`
}

func Init() {
	logfile := "log/tasks.log"
	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE, os.ModePerm)

	if err == nil {
		log.SetOutput(file)
		log.SetPrefix("task:")
		log.SetFlags(log.Flags() | log.Lshortfile)
	}
}

func loadTask() ([]Task, error) {
	if bytes, err := ioutil.ReadFile("datas/tasks.json"); err != nil {
		if os.IsNotExist(err) {
			log.Println("tasks.json 文件不存在")
			return []Task{}, nil
		}
		log.Fatalln("load Task error: ", err)
		return nil, err
	} else {
		var tasks []Task
		if err := json.Unmarshal(bytes, &tasks); err == nil {
			log.Println("task load susscess")
			return tasks, nil

		} else {
			log.Fatalln("load Task error: ", err)
			return nil, err
		}
	}
}

func storeTasks(tasks []Task) error {
	bytes, err := json.Marshal(tasks)
	if err != nil {
		log.Fatalln("store Task error: ", err)
		return nil
	}
	log.Println("task store susscess")
	return ioutil.WriteFile("datas/tasks.json", bytes, 0X066)
}

func GetTasks() []Task {
	tasks, err := loadTask()
	if err == nil {
		return tasks
	}
	panic(err)
}

func GetTaskId() (int, error) {
	tasks, err := loadTask()
	if err != nil {
		return -1, err
	}

	var id int
	for _, task := range tasks {
		if id < task.Id {
			id = task.Id
		}
	}

	return id + 1, nil
}

func CreateTask(name, user, desc string) {
	id, err := GetTaskId()
	if err != nil {
		log.Fatalln("task id load: ", err)
		panic(err)
	}
	task := Task{
		Id:       id,
		Name:     name,
		User:     user,
		Desc:     desc,
		Progress: 0,
		Status:   "new",
	}
	tasks, err := loadTask()
	if err != nil {
		log.Fatalln("load task err: ", err)
		panic(err)
	}
	tasks = append(tasks, task)
	log.Println("task create successful")
	storeTasks(tasks)
}

func GetTaskById(id int) (Task, error) {
	tasks, err := loadTask()
	if err != nil {
		panic(err)
	}
	for _, task := range tasks {
		if id == task.Id {
			return task, nil
		}
	}
	return Task{}, errors.New("Not Found")
}

func ModifyTask(id int, name, desc string, progress int, user, status string) {
	tasks, err := loadTask()
	if err != nil {
		panic(err)
	}

	newTasks := make([]Task, len(tasks))

	for i, task := range tasks {
		if id == task.Id {
			task.Name = name
			task.Desc = desc
			task.Progress = progress
			task.User = user
			task.Status = status
		}
		newTasks[i] = task
	}
	storeTasks(newTasks)
}

func DeleteTask(id int) {
	tasks, err := loadTask()
	if err != nil {
		panic(err)
	}

	newTasks := make([]Task, 0)
	for _, task := range tasks {
		if id != task.Id {
			newTasks = append(newTasks, task)
		}
	}
	storeTasks(newTasks)
}
