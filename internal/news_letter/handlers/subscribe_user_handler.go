package handlers

import (
	"go_newsletter_api/internal/auth"
	"go_newsletter_api/internal/news_letter/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SubscribeUserHandler(newsletterService *service.NewsletterService) gin.HandlerFunc {
	return func(c *gin.Context) {
		newsletterID, _ := strconv.ParseUint(c.Param("newsletterID"), 10, 32)

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

		userID := claims.UserID

		err = newsletterService.SubscribeUser(uint(newsletterID), uint(userID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User subscribed successfully"})
	}
}
