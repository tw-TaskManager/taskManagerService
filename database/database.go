package database

import (
	"database/sql"
	"taskManagerService/model"
	_"github.com/lib/pq"
)

func SaveTask(db *sql.DB, tasks *model.Task) (error) {
	_, queryErr := db.Exec("INSERT INTO Task_Manager (task) VALUES($1)", tasks.Task)
	if (queryErr != nil) {
		return queryErr;
	}
	return nil;
}

func GetTasks(db *sql.DB) ([]*model.Task, error) {
	rows, err := db.Query("SELECT id,task from Task_Manager")
	if (err != nil) {
		return nil, err;
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
