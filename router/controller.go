package router

import (
	"github.com/gorilla/mux"
	"strings"
	"net/http"
	"task_manager/database"
	"github.com/mholt/binding"
	"task_manager/model"
	"log"
	"fmt"
	"database/sql"
)

func FieldMap(db *sql.DB) binding.FieldMap {
	return binding.FieldMap{
		db: db,
	}
}

func HandleRequest(db *sql.DB) {
	handler := mux.NewRouter()
	handler.HandleFunc("/tasks", saveTask(db)).Methods("POST")
	handler.HandleFunc("/tasks", getAllTask(db)).Methods("GET")
	handler.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
	http.Handle("/", handler)

}

func saveTask(db *sql.DB) http.HandlerFunc {
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
func getAllTask(db *sql.DB) http.HandlerFunc {
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
