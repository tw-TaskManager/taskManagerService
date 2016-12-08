package main

import (
	"net/http"
	"TODO_Maker/router"
	"TODO_Maker/database"
	"log"
)

func main() {

	db, err := database.CreateDataBase()
	if (err != nil) {
		log.Fatal(err.Error())
	}
	router.HandleRequest(db);
	defer db.Close()
	http.ListenAndServe(":3000", nil)

}
