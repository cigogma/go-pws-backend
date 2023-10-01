package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pws-backend/api/controller"
	"github.com/pws-backend/bootstrap"
	"github.com/pws-backend/repository"
	"github.com/pws-backend/usecase"
	"gorm.io/gorm"
)

func NewFilesRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewFileRepository(db)
	pc := &controller.ProjectsController{
		ProjectUsecase: usecase.NewProjectUsecase(ur, timeout),
	}
	group.GET("/files", pc.Index)
	group.GET("/files/:projectId", pc.Details)
	group.PUT("/files/:projectId", pc.Update)
	group.POST("/files", pc.Create)
}
