package router

import (
	"github.com/gorilla/mux"
	"strings"
	"net/http"
	"task_manager/database"
	"task_manager/model"
	"log"
	"fmt"
	"database/sql"
)

func HandleRequest(db *sql.DB) {
	handler := mux.NewRouter()
	handler.HandleFunc("/tasks", SaveTask(db)).Methods("POST")
	handler.HandleFunc("/tasks", GetAllTask(db)).Methods("GET")
	handler.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
	http.Handle("/", handler)

}

func SaveTask(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm();
		task := strings.Join(req.Form["task"], "");
		task_to_db := model.Task{Task:task}
		_, err := database.SaveTask(db, &task_to_db)
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

