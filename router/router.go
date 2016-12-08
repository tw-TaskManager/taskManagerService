package router

import "net/http"
import (
	"github.com/gorilla/mux"
)

func HandleRequest() {
	handler := mux.NewRouter()
	handler.HandleFunc("/hello", PrintHelloWorld).Methods("GET")
	handler.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
	http.Handle("/", handler)

}

func PrintHelloWorld(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hello world"))
}