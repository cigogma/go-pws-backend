package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pws-backend/api/controller"
	"github.com/pws-backend/bootstrap"
	"github.com/pws-backend/usecase"
	"gorm.io/gorm"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	lc := &controller.AuthLoginController{
		LoginUsecase: usecase.NewLoginUsecase(db, timeout),
		Env:          env,
	}
	group.POST("/auth/login", lc.Login)
}
