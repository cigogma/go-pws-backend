package domain

import (
	"context"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model

	Name        string
	Description string
	OrderIndex  int `gorm:"default:0"`
	ThumbnailID uint
	IsActive    bool

	Images []ProjectImage
}

type ProjectImage struct {
	gorm.Model

	OrderIndex int `gorm:"default:1"`
	FileID     uint
	ProjectID  uint
}

type ProjectRepository interface {
	Create(c context.Context, project *Project) error
	FetchByUserID(c context.Context, userID string) ([]Project, error)
}

type ProjectUsecase interface {
	Create(c context.Context, project *Project) error
	FetchByUserID(c context.Context) ([]Project, error)
}
