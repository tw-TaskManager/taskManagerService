package router

import (
	"database/sql"
	"github.com/gorilla/mux"
	taskHandler "taskManagerService/handler"
	"net/http"
)

func HandleRequest(db *sql.DB) {
	handler := mux.NewRouter()
	handler.HandleFunc("/tasks/save", taskHandler.SaveTask(db)).Methods("POST")
	handler.HandleFunc("/task/delete", taskHandler.DeleteTask(db)).Methods("POST")
	handler.HandleFunc("/task/update", taskHandler.UpdateTask(db)).Methods("POST")
	handler.HandleFunc("/tasks", taskHandler.GetAllTask(db)).Methods("GET")
	handler.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
	http.Handle("/", handler)
}
