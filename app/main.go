package main

import (
	"net/http"
	"taskManagerService/router"
	"taskManagerService/database"
	"log"
)

func main() {

	db, err := database.OpenDatabase()
	if (err != nil) {
		log.Fatal(err.Error())
	}
	db.Ping()
	defer db.Close()
	router.HandleRequest(db);
	http.ListenAndServe(":3000", nil)
}