package handlers

import (
	"github.com/gin-gonic/gin"
	"go_newsletter_api/internal/news_letter/service"
	"net/http"
	"strconv"
)

func UnsubscribeUserHandler(newsletterService *service.NewsletterService) gin.HandlerFunc {
	return func(c *gin.Context) {
		newsletterID, _ := strconv.ParseUint(c.Param("newsletterID"), 10, 32)
		userID, _ := strconv.ParseUint(c.Param("userID"), 10, 32)

		err := newsletterService.UnsubscribeUser(uint(newsletterID), uint(userID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User unsubscribed successfully"})
	}
}
