package domain

import (
	"context"
	"mime/multipart"
	"time"

	"gorm.io/gorm"
)

type File struct {
	ID           uint           `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Key          string         `json:"key"`
	Bucket       string         `json:"bucket"`
	Mime         string         `json:"mime"`
	StorageClass string         `json:"storageClass"`
	Comment      string         `json:"comment"`
}

type FileRepository interface {
	Create(c context.Context, file *File) error
	Delete(c context.Context, file *File) error
	GetByID(c context.Context, fileId uint) (File, error)
	Fetch(c context.Context, limit int) ([]File, error)
}

type FileUsecase interface {
	Upload(c context.Context, file multipart.File, fileE *File) error
	Create(c context.Context, file *File) error
	Delete(c context.Context, file *File) error
	GetByID(c context.Context, fileId uint) (File, error)
	Fetch(c context.Context, limit int) ([]File, error)
}
