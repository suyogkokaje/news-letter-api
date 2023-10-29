package main

import (
	"go_newsletter_api/db"
	edition_model "go_newsletter_api/internal/edition/model"
	edition_repository "go_newsletter_api/internal/edition/repository"
	edition_service "go_newsletter_api/internal/edition/service"
	news_letter_model "go_newsletter_api/internal/news_letter/model"
	news_letter_repository "go_newsletter_api/internal/news_letter/repository"
	news_letter_service "go_newsletter_api/internal/news_letter/service"
	user_model "go_newsletter_api/internal/user/model"
	user_repository "go_newsletter_api/internal/user/repository"
	user_service "go_newsletter_api/internal/user/service"
	"go_newsletter_api/routes"
	"go_newsletter_api/scheduler"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

	dbInstance.AutoMigrate(&user_model.User{})
	dbInstance.AutoMigrate(&news_letter_model.Newsletter{})
	dbInstance.AutoMigrate(&news_letter_model.NewsletterSubscriber{})
	dbInstance.AutoMigrate(&edition_model.Edition{}) 

	userRepo := user_repository.UserRepository{DB: dbInstance}
	userService := user_service.NewUserService(userRepo)

	newsletterRepo := news_letter_repository.NewsletterRepository{DB: dbInstance}
	newsletterService := news_letter_service.NewNewsletterService(newsletterRepo)

	editionRepo := edition_repository.EditionRepository{DB: dbInstance}
	editionService := edition_service.NewEditionService(editionRepo)

	scheduler.StartEditionPublishScheduler(*editionService, *newsletterService)

	routes.SetupRouter(r, userService, newsletterService, editionService) 

	r.Run(":8080")
}
