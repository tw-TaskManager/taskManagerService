package handler

import (
	"task_manager/database"
	"log"
	"database/sql"
	"task_manager/model"
	"net/http"
	"fmt"
	"strings"
)

func SaveTask(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm();
		task := strings.Join(req.Form["task"], "");
		task_to_db := model.Task{Task:task}
		err := database.SaveTask(db, &task_to_db)
		if (err != nil) {
			log.Fatal(err.Error())
			res.Write([]byte("got error.."))
			return
		}
		res.Write([]byte("task has stored.."));
	}
}
func GetAllTask(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		data, err := database.GetTasks(db);
		if (err != nil) {
			log.Fatal(err)
			res.Write([]byte("got error.."))
			return
		}
		var tasks string;
		for _, each := range data {
			fmt.Println(each)
			tasks += each.Task + "<br/>"
		}
		res.Write([]byte(tasks));
	}
}
