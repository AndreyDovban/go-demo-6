package auth_test

import (
	"go-demo-6/internal/auth"
	"go-demo-6/internal/user"
	"testing"
)

type MockUserRepository struct{}

func (repo *MockUserRepository) Create(u *user.User) (*user.User, error) {
	return &user.User{
		Email: "a@a.ru",
	}, nil
}

func (repo *MockUserRepository) FindByEmail(email string) (*user.User, error) {
	return nil, nil
}

func TestRegisterSuccess(t *testing.T) {
	const initEmail = "a@a.ru"

	authService := auth.NewAuthService(&MockUserRepository{})
	email, err := authService.Register(initEmail, "111", "Andrey")
	if err != nil {
		t.Fatal(err)
	}
	if email != initEmail {
		t.Fatalf("Email %s do not math %s", email, initEmail)
	}
}
