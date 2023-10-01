package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pws-backend/api/middleware"
	"github.com/pws-backend/bootstrap"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	InitRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)
	NewFileRouter(env, timeout, db, publicRouter)
	NewFilesRouter(env, timeout, db, publicRouter)
	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewProfileRouter(env, timeout, db, protectedRouter)
}
