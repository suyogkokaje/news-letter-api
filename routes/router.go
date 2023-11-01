package routes

import (
	"go_newsletter_api/internal/auth"
	edition_handlers "go_newsletter_api/internal/edition/handlers"
	edition_service "go_newsletter_api/internal/edition/service"
	news_letter_handlers "go_newsletter_api/internal/news_letter/handlers"
	news_letter_service "go_newsletter_api/internal/news_letter/service"
	user_handlers "go_newsletter_api/internal/user/handlers"
	user_service "go_newsletter_api/internal/user/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, userService *user_service.UserService, newsletterService *news_letter_service.NewsletterService, editionService *edition_service.EditionService) {
	userRoutes := r.Group("/user")
	{
		userRoutes.POST("/signup", func(c *gin.Context) {
			user_handlers.SignUpHandler(c, userService)
		})

		userRoutes.POST("/login", func(c *gin.Context) {
			user_handlers.LoginHandler(c, userService)
		})
	}

	newsletterRoutes := r.Group("/newsletter")
	{
		newsletterRoutes.Use(auth.AdminAuthMiddleware())

		newsletterRoutes.POST("/create", news_letter_handlers.CreateNewsletterHandler(newsletterService))
		newsletterRoutes.GET("/subscribers", news_letter_handlers.GetSubscribersHandler(newsletterService))
		newsletterRoutes.DELETE("/delete/:id", news_letter_handlers.DeleteNewsletterHandler(newsletterService))
	}

	subscriptionRoutes := r.Group("/newsletter")
	{
		subscriptionRoutes.Use(auth.UserAuthMiddleware())

		subscriptionRoutes.POST("/subscribe/:newsletterID", news_letter_handlers.SubscribeUserHandler(newsletterService))
		subscriptionRoutes.POST("/unsubscribe/:newsletterID", news_letter_handlers.UnsubscribeUserHandler(newsletterService))
		subscriptionRoutes.GET("/subscriptions", user_handlers.GetUserSubscriptionsHandler(userService))
	}

	editionRoutes := r.Group("/edition")
	{
		editionHandler := edition_handlers.NewEditionHandler(*editionService)
		editionRoutes.POST("/create", auth.AdminAuthMiddleware(), edition_handlers.CreateEditionHandler(editionHandler))
		editionRoutes.PUT("/update/:id", auth.AdminAuthMiddleware(), edition_handlers.UpdateEditionHandler(editionHandler))
		editionRoutes.GET("/get/:id", edition_handlers.GetEditionByIDHandler(editionHandler))
		editionRoutes.GET("/get-by-newsletter/:newsletterID", edition_handlers.GetEditionsByNewsletterIDHandler(editionHandler))
		editionRoutes.DELETE("/delete/:id", auth.AdminAuthMiddleware(), edition_handlers.DeleteEditionHandler(editionHandler))
	}
}
