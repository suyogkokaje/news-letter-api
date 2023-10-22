package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go_newsletter_api/db"
	"go_newsletter_api/internal/user/model"
	"go_newsletter_api/internal/user/repository"
	"go_newsletter_api/internal/user/service"
	"go_newsletter_api/routes"
)

func main() {
	connectionStr := "user=root dbname=newsletter password=password host=localhost port=5000 sslmode=disable"

	db.Initialize(connectionStr)

	r := gin.Default()

	dbInstance, err := gorm.Open("postgres", connectionStr)
	if err != nil {
		panic("Failed to connect to the database")
	}
	defer dbInstance.Close()

	dbInstance.AutoMigrate(&model.User{})

	userRepo := repository.UserRepository{DB: dbInstance}
	userService := service.NewUserService(userRepo)

	routes.SetupRouter(r, userService)

	r.Run(":8080")
}
