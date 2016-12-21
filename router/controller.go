package router

import (
	"database/sql"
	"github.com/gorilla/mux"
	taskHandler "taskManagerService/handler"
	"net/http"
)

func HandleRequest(db *sql.DB) {
	handler := mux.NewRouter()
	handler.HandleFunc("/tasks/save/{id:[a-z]+}", taskHandler.SaveTask(db)).Methods("POST")
	handler.HandleFunc("/task/delete/{id:[a-z]+}", taskHandler.DeleteTask(db)).Methods("POST")
	handler.HandleFunc("/task/update/{id:[a-z]+}", taskHandler.UpdateTask(db)).Methods("POST")
	handler.HandleFunc("/tasks/{id:[a-z]+}", taskHandler.GetAllTask(db)).Methods("GET")
	handler.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
	http.Handle("/", handler)
}
