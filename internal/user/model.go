package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" validate:"required,email" grom:"index"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}
