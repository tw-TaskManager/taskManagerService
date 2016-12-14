package database

import (
	"database/sql"
	"taskManagerService/model"
	_"github.com/lib/pq"
	"encoding/json"
)

func SaveTask(db *sql.DB, tasks *model.Task) (error) {
	_, queryErr := db.Exec("INSERT INTO Task_Manager (task) VALUES($1)", tasks.Task)
	if (queryErr != nil) {
		return queryErr;
	}
	return nil;
}

func GetTasks(db *sql.DB) ([]byte, error) {
	rows, err := db.Query("SELECT id,task from Task_Manager")
	if (err != nil) {
		return nil, err;
	}

	listOfTasks := []*model.Task{}
	for rows.Next() {
		var task string
		var id int
		rows.Scan(&id, &task)
		tasks := model.Task{task, id}
		listOfTasks = append(listOfTasks, &tasks)
	}
	data, err := json.Marshal(listOfTasks)
	return data, nil
}


// DB Migration - goose - go library
