package main

import (
	"net/http"
	"TODO_Maker/router"
)

func main() {
	router.HandleRequest();
	http.ListenAndServe(":3000", nil)

}
