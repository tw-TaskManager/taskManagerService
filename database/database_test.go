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
	userId := fmt.Sprint("1")
	rows := sqlmock.NewRows([]string{"userId"}).
		AddRow(1)

	mock.ExpectQuery("INSERT INTO task_manager").WithArgs(task.Task, userId).WillReturnRows(rows)//assert
	id, err := SaveTask(db, &task, userId)
	assert.Nil(t, err)
	assert.Equal(t, id, int32(1))
}

func TestSaveTasksReturnError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db.Close();
	task := model.Task{Task:"this is task"}
	mock.ExpectExec("INSERT INTO Task_Manager").WithArgs(task.Task).WillReturnError(fmt.Errorf("database is closed")) //assert
	userId := fmt.Sprint("1")
	id, err := SaveTask(db, &task, userId);
	assert.NotNil(t, err)
	assert.Equal(t, id, int32(0))
}

func TestDeleteTaskForSuccessConnection(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close();
	taskId := model.Task{Task:"this is task", Id:2}
	userId := fmt.Sprint("2")
	mock.ExpectExec("DELETE FROM task_manager").
		WithArgs(taskId.Id, userId).WillReturnResult(sqlmock.NewResult(1, 1))
	err = DeleteTask(db, &taskId, userId)
	assert.Nil(t, err)
}

func TestDeleteTaskForDbError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	taskId := model.Task{Task:"this is task", Id:2}
	userId := fmt.Sprint("2")

	mock.ExpectExec("DELETE FROM Task_Manager").
		WithArgs(taskId.Id, userId).
		WillReturnError(err)

	err = DeleteTask(db, &taskId, userId)
	assert.NotNil(t, err)

}

func TestUpdateTaskForSuccessConnection(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close();
	tasks := model.Task{Task:"this is task", Id:2}
	mock.ExpectExec("UPDATE task_manager").WithArgs(tasks.Task, tasks.Id,"1").WillReturnResult(sqlmock.NewResult(1, 1))
	userId := fmt.Sprint("1")
	err = UpdateTask(db, &tasks, userId)
	assert.Nil(t, err)
}

func TestGetTasks(t *testing.T) {
	db, mock, err := sqlmock.New();

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err);
	}
	defer db.Close();
	rows := sqlmock.NewRows([]string{"id", "title"}).
		AddRow(1, "one")
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	userId := fmt.Sprint("1")
	task, err := GetTasks(db, userId);
	e := []*model.Task{}
	e = append(e, &model.Task{Task:"one", Id:1})
	assert.Equal(t, task, e)

}

func TestGetTasksWhenDBIsClosed(t *testing.T) {
	db, mock, err := sqlmock.New();

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err);
	}
	db.Close();
	mock.ExpectQuery("SELECT").WillReturnError(fmt.Errorf("database is closed"))
	userId := fmt.Sprint("1")
	rs, err := GetTasks(db, userId);
	assert.Equal(t, len(rs), 0)
	assert.Error(t, err);

}

