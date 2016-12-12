package database

import (
	"testing"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"task_manager/model"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestSaveTasksForSuccessConnection(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close();
	task := model.Task{Task:"this is task"}
	mock.ExpectExec("INSERT INTO Task_Manager").WithArgs(task.Task).WillReturnResult(sqlmock.NewResult(1, 2)) //assert
	if err := SaveTask(db, &task); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestSaveTasksReturnError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db.Close();
	task := model.Task{Task:"this is task"}
	mock.ExpectExec("INSERT INTO Task_Manager").WithArgs(task.Task).WillReturnError(fmt.Errorf("database is closed")) //assert
	err = SaveTask(db, &task);
	if err == nil {
		t.Errorf("was expecting an error, but there was none")
	}
}

func TestGetTasks(t *testing.T) {
	db, mock, err := sqlmock.New();

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err);
	}
	defer db.Close();
	rows := sqlmock.NewRows([]string{"id", "title"}).
		AddRow(1, "one").
		AddRow(2, "two")
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	rs, _ := GetTasks(db);
	assert.Equal(t, len(rs), 2)

}

func TestGetTasksWhenDBIsClosed(t *testing.T) {
	db, mock, err := sqlmock.New();

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err);
	}
	db.Close();
	mock.ExpectQuery("SELECT").WillReturnError(fmt.Errorf("database is closed"))
	rs, err := GetTasks(db);
	assert.Equal(t, len(rs), 0)
	assert.Error(t, err);

}

