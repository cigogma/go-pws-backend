package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID          uint           `json:"id" gorm:"primary_key"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	OrderIndex  int            `json:"orderIndex" gorm:"default:0"`
	ThumbnailID uint           `json:"thumbnailId"`
	IsActive    bool           `json:"isActive" gorm:"default:true"`
	Thumbnail   File           `json:"thumbnail" gorm:"foreignKey:ThumbnailID"`
	Images      []ProjectImage `json:"images" gorm:"foreignKey:ProjectID"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type ProjectImage struct {
	ID         uint           `json:"id" gorm:"primary_key"`
	OrderIndex int            `gorm:"default:1"`
	FileID     uint           `json:"fileId"`
	ProjectID  uint           `json:"projectId"`
	File       File           `json:"file" gorm:"foreignKey:FileID"`
	Project    Project        `json:"project" gorm:"foreignKey:ProjectID"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt" gorm:"index"`
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
