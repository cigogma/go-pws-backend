package bootstrap

import (
	"gorm.io/gorm"
)

type Application struct {
	Env *Env
	DB  *gorm.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.DB = NewDatabase(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseDBConnection(app.DB)
}
