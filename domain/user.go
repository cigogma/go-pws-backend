package domain

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context, limit int) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByID(c context.Context, id uint) (User, error)
}
