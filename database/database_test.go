package database

import (
	"testing"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"taskManagerService/model"
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
	if _, err := SaveTask(db, &task); err != nil {
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
	_,err = SaveTask(db, &task);
	if err == nil {
		t.Errorf("was expecting an error, but there was none")
	}
}

func TestDeleteTaskForSuccessConnection(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close();
	taskId := model.Task{Task:"this is task", Id:2}
	mock.ExpectExec("DELETE FROM Task_Manager").WithArgs(taskId.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	if err := DeleteTask(db, &taskId); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func TestUpdateTaskForSuccessConnection(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close();
	tasks := model.Task{Task:"this is task", Id:2}
	mock.ExpectExec("UPDATE Task_Manager SET").WithArgs(tasks.Task, tasks.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	if err := UpdateTask(db, &tasks); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
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
	GetTasks(db);
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}

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

