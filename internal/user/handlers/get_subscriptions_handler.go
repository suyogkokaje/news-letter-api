package handlers

import (
	"github.com/gin-gonic/gin"
	"go_newsletter_api/internal/auth"
	"go_newsletter_api/internal/user/service"
	"net/http"
)

func GetUserSubscriptionsHandler(userService *service.UserService) gin.HandlerFunc {
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

		userID := claims.UserID

		newsletters, err := userService.GetUserSubscriptions(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, newsletters)
	}
}
