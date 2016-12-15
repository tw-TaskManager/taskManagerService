package database

import (
	"database/sql"
	"taskManagerService/model"
	_"github.com/lib/pq"
	"encoding/json"
)

func SaveTask(db *sql.DB, tasks *model.Task) (int32, error) {
	result, queryErr := db.Query("INSERT INTO Task_Manager (task) VALUES($1) RETURNING id;", tasks.Task)

	if (queryErr != nil) {
		return 0, queryErr;
	}
	ids := make([]int32, 0, 0)
	for result.Next() {
		var id int32;
		result.Scan(&id)
		ids = append(ids, id)
	}
	currentId := ids[len(ids) - 1]
	return currentId, nil;
}

func GetTasks(db *sql.DB) ([]byte, error) {
	rows, err := db.Query("SELECT id,task from Task_Manager")
	if (err != nil) {
		return nil, err;
	}

	listOfTasks := []*model.Task{}
	for rows.Next() {
		var task string
		var id int32
		rows.Scan(&id, &task)
		tasks := model.Task{task, id}
		listOfTasks = append(listOfTasks, &tasks)
	}
	data, err := json.Marshal(listOfTasks)
	return data, nil
}

func DeleteTask(db *sql.DB, task *model.Task) (error) {
	_, queryErr := db.Exec("DELETE FROM Task_Manager WHERE id=$1", task.Id);
	if (queryErr != nil) {
		return queryErr;
	}
	return nil;
}

func UpdateTask(db *sql.DB, task *model.Task) (error) {
	_, queryErr := db.Exec("UPDATE Task_Manager SET task=$1 where id=$2;", task.Task, task.Id);
	if (queryErr != nil) {
		return queryErr;
	}
	return nil;
}


// DB Migration - goose - go library
