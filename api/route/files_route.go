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
	fr := repository.NewFileRepository(db)
	pc := &controller.FilesController{
		FileUsecase: usecase.NewFileUsecase(fr, timeout),
	}
	group.GET("/files", pc.Index)
	group.GET("/files/:fileId", pc.Details)
	group.DELETE("/files/:fileId", pc.Delete)
	group.POST("/files", pc.Create)
}
