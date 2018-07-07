package myapp

import "testing"

type mockService struct {
	UserService
}

func (s mockService) User(id int) (*User, error) {
	return &User{
		Name: "Odds",
	}, nil
}

func TestUserCacheGetUserBuID(t *testing.T) {
	service := &mockService{}
	uc := NewUserCache(service)
	u, err := uc.User(1)
	if err != nil {
		t.Fatal(err)
	}

	if u.Name != "Odds" {
		t.Fatalf("Got wrong data\n%#v\n", u)
	}
}
