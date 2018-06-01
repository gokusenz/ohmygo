package mock

import "github.com/pallat/ohmygo/myapp"

// UserService represents a mock implementation of myapp.UserService.
type UserService struct {
	UserFn      func(id int) (*myapp.User, error)
	UserInvoked bool

	UsersFn      func() ([]*myapp.User, error)
	UsersInvoked bool

	// additional function implementations...
}

// User invokes the mock implementation and marks the function as invoked.
func (s *UserService) User(id int) (*myapp.User, error) {
	s.UserInvoked = true
	return s.UserFn(id)
}

func (s *UserService) Users() ([]*User, error) {
	return nil, nil
}
func (s *UserService) CreateUser(u *User) error {
	return nil
}
func (s *UserService) DeleteUser(id int) error {
	return nil
}
