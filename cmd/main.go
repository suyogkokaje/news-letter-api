package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go_newsletter_api/db"
	news_letter_repository "go_newsletter_api/internal/news_letter/repository"
	news_letter_service "go_newsletter_api/internal/news_letter/service"
	news_letter_model "go_newsletter_api/internal/news_letter/model"
	user_model "go_newsletter_api/internal/user/model"
	user_repository "go_newsletter_api/internal/user/repository"
	user_service "go_newsletter_api/internal/user/service"
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

	dbInstance.AutoMigrate(&user_model.User{}, &news_letter_model.Newsletter{})
	userRepo := user_repository.UserRepository{DB: dbInstance}
	userService := user_service.NewUserService(userRepo)

	newsletterRepo := news_letter_repository.NewsletterRepository{DB: dbInstance}
	newsletterService := news_letter_service.NewNewsletterService(newsletterRepo)

	routes.SetupRouter(r, userService, newsletterService)

	r.Run(":8080")
}
