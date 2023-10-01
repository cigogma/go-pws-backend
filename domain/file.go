package domain

import (
	"context"

	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	ID           uint `json:"id" gorm:"primary_key"`
	Key          string
	Bucket       string
	Mime         string
	StorageClass string
	Comment      string
}

type FileRepository interface {
	Create(c context.Context, project *File) error
	GetByID(c context.Context, projectId uint) (*File, error)
	Fetch(c context.Context, limit int) ([]File, error)
}

type FileUsecase interface {
	Create(c context.Context, project *File) error
	GetByID(c context.Context, projectId uint) (*File, error)
	Fetch(c context.Context, limit int) ([]File, error)
}
