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
	_"encoding/json"
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
		return
	}
}
func GetAllTask(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		responseContract := contract.Response{}
		data, err := database.GetTasks(db);
		responseContract.Response = []byte(data)
		if (err != nil) {
			res.Write([]byte("got error."))
			return
		}
		dataToSend, err := proto.Marshal(&responseContract)
		if (err != nil) {
			res.Write([]byte("got error."))
			return
		}
		res.Write(dataToSend)
	}
}
