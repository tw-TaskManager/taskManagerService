package database

import (
	"database/sql"
	"taskManagerService/model"
	_"github.com/lib/pq"
	"fmt"
)

const (
	InsertTaskOfGivenId = `INSERT INTO task_manager (task,userId) VALUES($1,$2) RETURNING id;`
	UpdateTaskOfGivenId = `UPDATE task_manager SET task=$1 where id=$2 and userId=$3;`
	DeleteTaskOfGivenId = `DELETE FROM task_manager WHERE id=$1 and userId=$2`
	SelectTaskOfGivenId = `SELECT id,task from task_manager where userId=$1`
)

func SaveTask(db *sql.DB, tasks *model.Task, userId string) (int32, error) {
	result, queryErr := db.Query(InsertTaskOfGivenId, tasks.Task, userId)
	if (queryErr != nil) {
		fmt.Println("----------------")
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

func GetTasks(db *sql.DB, userId string) ([]*model.Task, error) {
	rows, err := db.Query(SelectTaskOfGivenId, userId)
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
	return listOfTasks, nil
}

func DeleteTask(db *sql.DB, task *model.Task, userId string) (error) {
	_, queryErr := db.Exec(DeleteTaskOfGivenId, task.Id, userId);
	if (queryErr != nil) {
		return queryErr;
	}
	return nil;
}

func UpdateTask(db *sql.DB, task *model.Task, userId string) (error) {
	_, queryErr := db.Exec(UpdateTaskOfGivenId, task.Task, task.Id, userId);
	if (queryErr != nil) {
		return queryErr;
	}
	return nil;
}


// DB Migration - goose - go library
