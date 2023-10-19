package repository

import (
	"context"

	"github.com/pws-backend/domain"
	"gorm.io/gorm"
)

type projectRepository struct {
	database *gorm.DB
}

func NewProjectRepository(db *gorm.DB) domain.ProjectRepository {
	return &projectRepository{
		database: db,
	}
}

func (ur *projectRepository) Create(c context.Context, project *domain.Project) error {
	err := ur.database.Create(project).Error
	return err
}

func (ur *projectRepository) Fetch(c context.Context, limit int) ([]domain.Project, error) {
	var users []domain.Project
	err := ur.database.Limit(limit).Preload("Thumbnail").Find(&users).Error
	return users, err
}

func (ur *projectRepository) GetByID(c context.Context, id uint) (*domain.Project, error) {
	var project *domain.Project
	err := ur.database.Model(domain.User{ID: id}).First(project).Error
	return project, err
}
