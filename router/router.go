package router

import (
	"github.com/gorilla/mux"
	"strings"
	"net/http"
	"toMaker/database"
	"toMaker/model"
	"log"
)

func HandleRequest() {
	handler := mux.NewRouter()
	handler.HandleFunc("/save", SaveTask).Methods("POST")
	handler.HandleFunc("/getData", GetTask).Methods("GET")
	handler.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
	http.Handle("/", handler)

}

func SaveTask(res http.ResponseWriter, req *http.Request) {
	req.ParseForm();
	task := strings.Join(req.Form["task"], "");
	task_to_db := model.Tasks{Id:7, Task:task}
	_, err := database.SaveTask(&task_to_db)
	if (err != nil) {
		log.Fatal(err)
		res.Write([]byte("got error.."))
	}
	res.Write([]byte("task has stored.."));
}

func GetTask(res http.ResponseWriter, req *http.Request) {
	data, err := database.GetTask();
	if (err != nil) {
		log.Fatal(err)
		res.Write([]byte("got error.."))
	}
	var tasks string;

	for _, each := range data {
		tasks += each.Task + "\n"
	}
	res.Write([]byte(tasks));
}
