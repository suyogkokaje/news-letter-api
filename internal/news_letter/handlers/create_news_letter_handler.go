package handlers

import (
	// "fmt"
	"go_newsletter_api/internal/auth"
	"go_newsletter_api/internal/news_letter/model"
	"go_newsletter_api/internal/news_letter/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateNewsletterHandler(newsletterService *service.NewsletterService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newsletter model.Newsletter

		if err := c.ShouldBindJSON(&newsletter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Retrieve the auth_token cookie
		token, err := c.Cookie("auth_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Parse and validate the token
		claims, err := auth.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Check if the user role is "admin" (or any other authorization checks you need)
		if claims.Role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}

		adminID := claims.UserID

		createdNewsletter, err := newsletterService.CreateNewsletter(&newsletter, adminID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, createdNewsletter)
	}
}
