package database

import (
	"database/sql"
	"toMaker/model"
	_"github.com/lib/pq"
)

const (
	DB_DRIVER = "postgres"
	DB_CONNECTION = "user=postgres dbname=postgres password=postgres sslmode=disable"
)

func GetDatabase() (*sql.DB, error) {
	db, err := sql.Open(DB_DRIVER, DB_CONNECTION)
	return db, err;
}

func SaveContact(c *model.Contact) (int, error) {
	db, err := GetDatabase();
	defer db.Close();
	if (err != nil) {
		return 0, err;
	}
	var id int;
	row := db.QueryRow("INSERT INTO CONTACT (id,name) VALUES($1,$2);", c.Id, c.Name).Scan(&id);
	if (row != nil) {
		return 0, err;
	}
	return id, nil;
}

func GetName() ([]*model.Contact, error) {
	db, err := GetDatabase();
	defer db.Close();
	if (err != nil) {
		return [] *model.Contact{}, err
	}
	rows, err := db.Query("SELECT * from contact")
	if (err != nil) {
		return [] *model.Contact{}, err;
	}
	var contacts []*model.Contact
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		contacts = append(contacts, &model.Contact{Id:id, Name:name})
	}
	return contacts, nil
}