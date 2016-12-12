package database

import (
	"database/sql"
	"task_manager/model"
	_"github.com/lib/pq"
	"log"
)

func SaveTask(db *sql.DB, tasks *model.Task) (int, error) {
	var id int;
	queryErr := db.QueryRow("INSERT INTO Task_Manager (task) VALUES($1) returning id;", tasks.Task).Scan(&id);
	if (queryErr != nil) {
		log.Fatal(queryErr)
		return 0, queryErr;
	}
	return id, nil;
}

func GetTasks(db *sql.DB) ([]*model.Task, error) {
	rows, err := db.Query("SELECT id,task from Task_Manager")
	if (err != nil) {
		log.Fatal(err.Error())
		return [] *model.Task{}, err;
	}
	var tasks []*model.Task
	for rows.Next() {
		var task string
		var id int
		rows.Scan(&id, &task)
		tasks = append(tasks, &model.Task{Task:task})
	}
	return tasks, nil
}


// DB Migration - goose - go library
