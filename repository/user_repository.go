package repository

import (
	"context"

	"github.com/pws-backend/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		database: db,
	}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	err := ur.database.Create(user).Error
	return err
}

func (ur *userRepository) Fetch(c context.Context, limit int) ([]domain.User, error) {
	var users []domain.User
	err := ur.database.Limit(limit).Find(&users).Error
	return users, err
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	var user domain.User
	err := ur.database.Model(domain.User{Email: email}).First(&user).Error
	return user, err
}

func (ur *userRepository) GetByID(c context.Context, id uint) (domain.User, error) {
	var user domain.User
	err := ur.database.Model(domain.User{ID: id}).First(&user).Error
	return user, err
}
