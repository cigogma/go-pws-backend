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

func NewProjectsRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewProjectRepository(db)
	pc := &controller.ProjectsController{
		ProjectUsecase: usecase.NewProjectUsecase(ur, timeout),
	}
	//sadas
	group.GET("/projects", pc.Index)
	group.GET("/projects/:projectId", pc.Details)
	group.PUT("/projects/:projectId", pc.Update)
	group.POST("/projects", pc.Create)
}
