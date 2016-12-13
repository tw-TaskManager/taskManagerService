package handler

import (
	"taskManagerService/database"
	"log"
	"database/sql"
	"net/http"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"taskManagerClient/contract"
	"taskManagerService/model"
)

func SaveTask(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm();
		requestData, err := ioutil.ReadAll(req.Body)
		if (err != nil) {
			log.Fatalf("got error while reading req %s", req.URL)
			return
		}
		data := &contract.Task{}

		if err = proto.Unmarshal(requestData, data); err != nil {
			log.Fatalln("got error while unmarsling")
		}
		taskToDb := model.Task{}
		taskToDb.Task = *data.Task

		if err = database.SaveTask(db, &taskToDb); err != nil {
			log.Fatal(err.Error())
			res.Write([]byte("got error.."))
			return
		}
		res.Write([]byte("task has stored.."));
	}
}
func GetAllTask(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		data, err := database.GetTasks(db);
		if (err != nil) {
			res.Write([]byte("got error.."))
			return
		}
		var tasks string;
		for _, each := range data {
			tasks += each.Task + "<br/>"
		}
		res.Write([]byte(tasks));
	}
}
