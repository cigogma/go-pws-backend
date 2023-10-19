package repository

import (
	"context"

	"github.com/pws-backend/domain"
	"gorm.io/gorm"
)

type fileRepository struct {
	database *gorm.DB
}

func NewFileRepository(db *gorm.DB) domain.FileRepository {
	return &fileRepository{
		database: db,
	}
}

func (ur *fileRepository) Create(c context.Context, file *domain.File) error {
	err := ur.database.Create(file).Error
	return err
}

func (ur *fileRepository) Delete(c context.Context, file *domain.File) error {
	return ur.database.Delete(file).Error
}
func (ur *fileRepository) Fetch(c context.Context, limit int) ([]domain.File, error) {
	var files []domain.File
	err := ur.database.Limit(limit).Find(&files).Error
	return files, err
}

func (ur *fileRepository) GetByID(c context.Context, id uint) (domain.File, error) {
	var file domain.File
	err := ur.database.First(&file, id).Error
	return file, err
}
