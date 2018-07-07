package postgres

import (
	"fmt"
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println("error creating mock database")
		return
	}

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "Stub Name 1")
	sql := `SELECT id, name FROM users WHERE id=?`

	// Stub/Mocking
	mock.ExpectQuery(sql).WillReturnRows(rows)

	u := UserService{
		db,
		nil,
	}

	result, err := u.User(1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n", result)
}
