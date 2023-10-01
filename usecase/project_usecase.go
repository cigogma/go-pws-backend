package usecase

import (
	"context"
	"time"

	"github.com/pws-backend/domain"
)

type projectUsecase struct {
	projectRepository domain.ProjectRepository
	contextTimeout    time.Duration
}

func NewProjectUsecase(projectRepository domain.ProjectRepository, timeout time.Duration) domain.ProjectUsecase {
	return &projectUsecase{
		projectRepository: projectRepository,
		contextTimeout:    timeout,
	}
}

func (pu *projectUsecase) Create(c context.Context, project *domain.Project) error {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.projectRepository.Create(ctx, project)
}

func (pu *projectUsecase) Fetch(c context.Context, limit int) ([]domain.Project, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.projectRepository.Fetch(ctx, limit)
}

func (pu *projectUsecase) GetByID(c context.Context, id uint) (*domain.Project, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()
	return pu.projectRepository.GetByID(ctx, id)
}
