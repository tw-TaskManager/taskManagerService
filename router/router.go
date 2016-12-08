package router

import (
	"github.com/gorilla/mux"
	"strings"
	"database/sql"
	"fmt"
	"net/http"
)

func HandleRequest() {
	handler := mux.NewRouter()
	handler.HandleFunc("/hello", PrintHelloWorld).Methods("GET")
	handler.HandleFunc("/save", SaveData).Methods("POST")
	handler.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
	http.Handle("/", handler)

}

func PrintHelloWorld(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hello world"))
}

const (
	DB_USER = "postgres"
	DB_NAME = "hello"
)

func SaveData(res http.ResponseWriter, req *http.Request) {
	req.ParseForm();
	a := strings.Join(req.Form["data"], "");
	dbinfo := fmt.Sprintf("user=%s dbname=%s", DB_USER, DB_NAME);
	db, err := sql.Open("postgres", dbinfo)
	err = db.Ping();
	if (err != nil) {
		fmt.Printf(" 1 =======errr is==> %s ", err)
	}
	/// inserting data...
	//var lastInsertId int
	rows, err := db.Query("SELECT *from data;")
	err = db.Ping();
	if (err != nil) {
		fmt.Printf("errr is============> %s ", err)
	}
	println(rows)

	//defer db.Close()
	//fmt.Println("last inserted id =", lastInsertId)
	//rows, err := db.Query("SELECT * FROM userinfo")
	//checkErr(err)
	//
	//for rows.Next() {
	//	var uid int
	//	var username string
	//	var department string
	//	var created time.Time
	//	err = rows.Scan(&uid, &username, &department, &created)
	//	checkErr(err)
	//	fmt.Println("uid | username | department | created ")
	//	fmt.Printf("%3v | %8v | %6v | %6v\n", uid, username, department, created)
	//}

	res.Write([]byte(a));
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}