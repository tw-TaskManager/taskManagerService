package database

import (
	"database/sql"
	"TODO_Maker/model"
	_"github.com/lib/pq"
	"log"
)

const (
	DB_DRIVER = "postgres"
	DB_CONNECTION = "user=postgres dbname=postgres password=postgres sslmode=disable"
)

func GetDatabase() (*sql.DB, error) {
	db, err := sql.Open(DB_DRIVER, DB_CONNECTION)
	return db, err;
}

func SaveTask(tasks *model.Tasks) (int, error) {
	db, err := GetDatabase();
	defer db.Close();
	if (err != nil) {
		return 0, err;
	}
	var id int;
	queryErr := db.QueryRow("INSERT INTO Task_Manager (id,task) VALUES($1,$2) returning id;", tasks.Id, tasks.Task).Scan(&id);
	if (queryErr != nil) {
		log.Fatal(queryErr)
		return 0, queryErr;
	}
	return id, nil;
}

func GetTask() ([]*model.Tasks, error) {
	db, err := GetDatabase();
	defer db.Close();
	if (err != nil) {
		return [] *model.Tasks{}, err
	}
	rows, err := db.Query("SELECT * from Task_Manager")
	if (err != nil) {
		return [] *model.Tasks{}, err;
	}
	var tasks []*model.Tasks
	for rows.Next() {
		var id int
		var task string
		rows.Scan(&id, &task)
		tasks = append(tasks, &model.Tasks{Id:id, Task:task})
	}
	return tasks, nil
}