package handlers

import (
	"go_newsletter_api/internal/auth"
	"go_newsletter_api/internal/news_letter/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSubscribedNewslettersHandler(newsletterService *service.NewsletterService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("auth_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		claims, err := auth.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		if claims.Role != "user" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}

		userID := claims.UserID

		subscribedNewsletters, err := newsletterService.GetNewslettersForSubscriber(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, subscribedNewsletters)
	}
}
