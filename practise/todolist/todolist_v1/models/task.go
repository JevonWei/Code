package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Task struct {
	Id           int        `json:"id"`
	Name         string     `json:"name"`
	Progress     int        `json:"progress"`
	User         string     `json:"user"`
	Desc         string     `json:"desc"`
	Status       int        `json:"status"`
	CreateTime   *time.Time `json:"create_time"`
	CompleteTime *time.Time `json:"complete_time"`
}

func Init() {
	logFile := "logs/tasks.log"
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err == nil {
		log.SetOutput(file)
		log.SetPrefix("task:")
		log.SetFlags(log.Flags() | log.Lshortfile)
	} else {
		fmt.Println(err)
	}

}

func loadTask() ([]Task, error) {
	if bytes, err := ioutil.ReadFile("datas/tasks.json"); err != nil {
		if os.IsNotExist(err) {
			log.Println("tasks.json 文件不存在")
			return []Task{}, nil
		}
		log.Fatalln("Task Load error: ", err)
		return nil, err
	} else {
		if len(bytes) == 0 {
			log.Println("tasks.json数据为空")
			return []Task{}, nil
		}

		var tasks []Task
		if err := json.Unmarshal(bytes, &tasks); err == nil {
			log.Println("Task load Success")

			return tasks, nil
		} else {
			log.Fatalln("Task load error:", err)

			return nil, err
		}
	}
}

func storeTask(tasks []Task) error {
	bytes, err := json.Marshal(tasks)
	if err != nil {
		log.Fatalln("Task store error:", err)
		return err
	}

	log.Println("Task store Success")
	return ioutil.WriteFile("datas/tasks.json", bytes, 0X066)
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

	rows, err := db.Query("select id,name,progress,user,`desc`,status,create_time,complate_time from task")
	if err != nil {
		panic(err)
	}

	tasks := make([]Task, 0)
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.Id, &task.Name, &task.Progress, &task.User, &task.Desc, &task.Status, &task.CreateTime, &task.CompleteTime); err == nil {
			tasks = append(tasks, task)
		} else {
			fmt.Println(err)
		}
	}
	// fmt.Println(tasks)
	return tasks
}

func GetUserTasks(name string) []Task {
	tasks, err := loadTask()
	if err != nil {
		log.Panicln("Task Read Faild", err)
		panic(err)
		return []Task{}
	}

	tasksnew := make([]Task, 0)

	for _, task := range tasks {
		if name == task.Name {
			tasksnew = append(tasksnew, task)
		}
	}
	return tasksnew
}

// func GetTaskId() (int, error) {
// 	tasks, err := loadTask()
// 	if err != nil {
// 		log.Fatalln("GetTaskId error:", err)
// 		return -1, err
// 	}

// 	var id int
// 	for _, task := range tasks {
// 		if id < task.Id {
// 			id = task.Id
// 		}
// 	}

// 	return id + 1, nil
// }

func CreateTask(name, user, desc string) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	defer db.Close()

	_, err = db.Exec("insert into task(name, `desc`, user, create_time) values(?,?,?,?)", name, desc, user, time.Now())
	if err != nil {
		panic(err)
	}
	// id, err := GetTaskId()
	// if err != nil {
	// 	log.Panic(err)
	// }

	// task := Task{
	// 	Id:       id,
	// 	Name:     name,
	// 	User:     user,
	// 	Desc:     desc,
	// 	Progress: 0,
	// 	Status:   0,
	// }

	// if tasks, err := loadTask(); err == nil {
	// 	tasks = append(tasks, task)
	// 	log.Printf("%s Task Create successful\n", task.Name)
	// 	storeTask(tasks)
	// }

}

func GetTaskById(id int) (Task, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	defer db.Close()

	var task Task
	row := db.QueryRow("select id, name, progress, user, `desc`, status from task where id = ?", id)

	err = row.Scan(&task.Id, &task.Name, &task.Progress, &task.User, &task.Desc, &task.Status)

	return task, err

	// tasks, err := loadTask()
	// if err != nil {
	// 	log.Panic(err)
	// }

	// for _, task := range tasks {
	// 	if id == task.Id {
	// 		log.Printf("Modify Task Name为: %s\n", task.Name)
	// 		return task, nil
	// 	}
	// }
	// log.Fatalf("Id 为 %d Task Not Found\n", id)
	// return Task{}, errors.New("Task Not Found")
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

	_, err = db.Exec("update task set name=?, `desc`=?, progress=?, user=?, status=? where id =?",
		name, desc, progress, user, status, id)
	if err != nil {
		panic(err)
	}

	// tasks, err := loadTask()
	// if err != nil {
	// 	panic(err)
	// }

	// newTasks := make([]Task, len(tasks))
	// for i, task := range tasks {
	// 	if id == task.Id {
	// 		task.Name = name
	// 		task.Desc = desc
	// 		task.Progress = progress
	// 		task.User = user
	// 		task.Status = status
	// 		log.Printf("%s Task 已更改\n", task.Name)
	// 	}

	// 	newTasks[i] = task
	// }
	// storeTask(newTasks)
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

	_, err = db.Exec("delete from task where id = ?", id)
	if err != nil {
		panic(err)
	}
	// tasks, err := loadTask()
	// if err != nil {
	// 	log.Panic(err)
	// }

	// newTasks := make([]Task, 0)
	// for _, task := range tasks {
	// 	if id != task.Id {
	// 		newTasks = append(newTasks, task)
	// 	}
	// }
	// log.Printf("Id为%d的Task删除成功\n", id)
	// storeTask(newTasks)
}
