package routes

import (
	"go_newsletter_api/internal/auth"
	news_letter_handlers "go_newsletter_api/internal/news_letter/handlers"
	news_letter_service "go_newsletter_api/internal/news_letter/service"
	user_handlers "go_newsletter_api/internal/user/handlers"
	user_service "go_newsletter_api/internal/user/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, userService *user_service.UserService, newsletterService *news_letter_service.NewsletterService) {
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
	}

	subscriptionRoutes := r.Group("/newsletter")
	{
		subscriptionRoutes.Use(auth.UserAuthMiddleware())

		subscriptionRoutes.POST("/subscribe/:newsletterID", news_letter_handlers.SubscribeUserHandler(newsletterService))
		subscriptionRoutes.POST("/unsubscribe/:newsletterID", news_letter_handlers.UnsubscribeUserHandler(newsletterService))
	}
}
