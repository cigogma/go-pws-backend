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

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}
