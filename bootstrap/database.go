package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pws-backend/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(env *Env) *gorm.DB {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass
	dbSchema := env.DBSchema

	databaseURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbSchema)

	if dbUser == "" || dbPass == "" {
		log.Fatal("Database user or password not set.")
	}

	client, err := gorm.Open(mysql.Open(databaseURI), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	client.AutoMigrate(&domain.User{})
	client.AutoMigrate(&domain.Project{})
	return client
}

func CloseDBConnection(client *gorm.DB) {
	if client == nil {
		return
	}
	sqlDB, err := client.DB()
	if err == nil {
		return
	}
	sqlDB.Close()

	log.Println("Connection to mySql closed.")
}
