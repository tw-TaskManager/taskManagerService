package main

import (
	"toMaker/abc.TODOMaker.com/router"

	"net/http"
)

func main() {
	router.HandleRequest();
	http.ListenAndServe(":3000", nil)

}
