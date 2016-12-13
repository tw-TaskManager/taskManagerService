package handler

import (
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"net/http/httptest"
	"net/http"
	"fmt"
)

func TestGetAllTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	if (err != nil) {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	getHandler := GetAllTask(db)

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	rr := httptest.NewRecorder()
	getHandler.ServeHTTP(rr, req)
	fmt.Println(rr.Code)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
	if rr.Code != http.StatusOK {
		t.Errorf("server is not responding %d", rr.Code)
	}
}

func TestSaveTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	if (err != nil) {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	getHandler := GetAllTask(db)

	req := httptest.NewRequest(http.MethodPost, "/tasks", nil)
	rr := httptest.NewRecorder()
	getHandler.ServeHTTP(rr, req)
	fmt.Println(rr.Code)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
	if rr.Code != http.StatusOK {
		t.Errorf("server is not responding %d", rr.Code)
	}
}
