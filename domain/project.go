package domain

import (
	"context"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	ID          uint `json:"id" gorm:"primary_key"`
	Name        string
	Description string
	OrderIndex  int `gorm:"default:0"`
	ThumbnailID uint
	IsActive    bool           `gorm:"default:true"`
	Thumbnail   []File         `gorm:"foreignKey:ThumbnailID"`
	Images      []ProjectImage `gorm:"foreignKey:ProjectID"`
}

type ProjectImage struct {
	gorm.Model
	ID         uint `json:"id" gorm:"primary_key"`
	OrderIndex int  `gorm:"default:1"`
	FileID     uint
	ProjectID  uint
	File       File    `gorm:"foreignKey:FileID"`
	Project    Project `gorm:"foreignKey:ProjectID"`
}

type ProjectRepository interface {
	Create(c context.Context, project *Project) error
	FetchByUserID(c context.Context, userID string) ([]Project, error)
}

type ProjectUsecase interface {
	Create(c context.Context, project *Project) error
	FetchByUserID(c context.Context) ([]Project, error)
}
