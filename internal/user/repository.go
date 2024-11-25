package user

import "go-demo-6/pkg/db"

type UserRepository struct {
	*db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{
		Db: database,
	}
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.Db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) FindByEmail(email string) (*User, error) {
	var user User
	result := repo.Db.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
