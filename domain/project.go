package domain

import (
	"context"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	ID          uint           `json:"id" gorm:"primary_key"`
	Name        string         `json:"name"`
	Description string         `json:"orderIndex"`
	OrderIndex  int            `gorm:"default:0"`
	ThumbnailID uint           `json:"thumbnailId"`
	IsActive    bool           `gorm:"default:true"`
	Thumbnail   File           `gorm:"foreignKey:ThumbnailID"`
	Images      []ProjectImage `gorm:"foreignKey:ProjectID"`
}

type ProjectImage struct {
	gorm.Model
	ID         uint    `json:"id" gorm:"primary_key"`
	OrderIndex int     `gorm:"default:1"`
	FileID     uint    `json:"fileId"`
	ProjectID  uint    `json:"projectId"`
	File       File    `gorm:"foreignKey:FileID"`
	Project    Project `gorm:"foreignKey:ProjectID"`
}

type ProjectRepository interface {
	Create(c context.Context, project *Project) error
	GetByID(c context.Context, projectId uint) (*Project, error)
	Fetch(c context.Context, limit int) ([]Project, error)
}

type ProjectUsecase interface {
	Create(c context.Context, project *Project) error
	GetByID(c context.Context, projectId uint) (*Project, error)
	Fetch(c context.Context, limit int) ([]Project, error)
}
