package main

import (

	"net/http"
	"toMaker/router"
)

func main() {
	router.HandleRequest();
	http.ListenAndServe(":3000", nil)

}
