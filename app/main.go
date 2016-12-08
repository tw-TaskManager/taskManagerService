package main

import (
	"net/http"
	"TODO_Maker/router"
	"TODO_Maker/database"
	"log"
)

func main() {

	db, err := database.OreateDataBase()
	if (err != nil) {
		log.Fatal(err.Error())
	}
	db.Ping()
	defer db.Close()
	router.HandleRequest(db);
	http.ListenAndServe(":3000", nil)
}
