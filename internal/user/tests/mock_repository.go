package user_test

import (
	"fmt"
	user_pkg "userlist-api-test/internal/user"
)

type mockRepository struct{}

func (m *mockRepository) FindAll() ([]user_pkg.User, error) {
	return []user_pkg.User{
		{Name: "John", Email: "john@example.com"},
	}, nil
}

func (m *mockRepository) Create(user *user_pkg.User) error {
	if user.Name == "" {
		return fmt.Errorf("Name cannot be empty")
	}
	user.ID = 1
	return nil
}

func (m *mockRepository) Update(user *user_pkg.User) error {
	return nil
}

func (m *mockRepository) Delete(id uint) error {
	return nil
}
