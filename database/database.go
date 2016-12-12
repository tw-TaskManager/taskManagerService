package database

import (
	"database/sql"
	"task_manager/model"
	_"github.com/lib/pq"
)

func SaveTask(db *sql.DB, tasks *model.Tasks) (error) {
	_, queryErr := db.Exec("INSERT INTO Task_Manager (task) VALUES($1)", tasks.Task)
	if (queryErr != nil) {
		return queryErr;
	}
	return nil;
}

func GetTasks(db *sql.DB) ([]*model.Tasks, error) {
	rows, err := db.Query("SELECT id,task from Task_Manager")
	if (err != nil) {
		return [] *model.Tasks{}, err;
	}
	var tasks []*model.Tasks
	for rows.Next() {
		var task string
		var id int
		rows.Scan(&id, &task)
		tasks = append(tasks, &model.Tasks{Task:task})
	}
	return tasks, nil
}


// DB Migration - goose - go library
