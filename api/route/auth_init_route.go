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

func InitRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	sc := controller.AuthInitController{
		SignupUsecase: usecase.AuthInitUseCase(ur, timeout),
		Env:           env,
	}
	group.POST("/auth/init", sc.Init)
}
