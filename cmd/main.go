package main

import (
	"time"

	"github.com/gin-gonic/gin"
	route "github.com/pws-backend/api/route"
	"github.com/pws-backend/bootstrap"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	db := app.DB
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, db, gin)

	gin.Run(env.ServerAddress)
}
