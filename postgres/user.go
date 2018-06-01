package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/pallat/ohmygo/myapp"
)

// UserService represents a PostgreSQL implementation of myapp.UserService.
type UserService struct {
	DB *sql.DB
}

func Open(s string) (*sql.DB, error) {
	return &sql.DB{}, nil
}

// User returns a user for a given id.
func (s *UserService) User(id int) (*myapp.User, error) {
	var u myapp.User
	row := s.DB.QueryRow(`SELECT id, name FROM users WHERE id = $1`, id)
	if err := row.Scan(&u.ID, &u.Name); err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *UserService) Users() ([]*myapp.User, error) {
	return nil, nil
}
func (s *UserService) CreateUser(u *myapp.User) error {
	return nil
}
func (s *UserService) DeleteUser(id int) error {
	return nil
}
