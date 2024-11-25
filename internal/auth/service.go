package auth

import (
	"errors"
	"go-demo-6/internal/user"
	"go-demo-6/pkg/di"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository di.IUserRepository
}

func NewAuthService(userRepository di.IUserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (service *AuthService) Register(email, password, name string) (string, error) {

	existedUser, _ := service.UserRepository.FindByEmail(email)
	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := &user.User{
		Email:    email,
		Password: string(hashPassword),
		Name:     name,
	}

	_, err = service.UserRepository.Create(user)
	if err != nil {
		return "", err
	}
	return user.Email, nil
}

func (service *AuthService) Login(email, password string) (string, error) {

	user, _ := service.UserRepository.FindByEmail(email)
	if user == nil {
		return "", errors.New(ErrWrongCredentials)
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New(ErrWrongCredentials)
	}

	return user.Email, nil
}
