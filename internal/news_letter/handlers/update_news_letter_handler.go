package handlers

import (
	"go_newsletter_api/internal/auth"
	"go_newsletter_api/internal/news_letter/model"
	"go_newsletter_api/internal/news_letter/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateNewsletterHandler(newsletterService *service.NewsletterService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newsletter model.Newsletter

		if err := c.ShouldBindJSON(&newsletter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

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

		if claims.Role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}

		newsletterID := c.Param("id")
		if newsletterID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Newsletter ID is required"})
			return
		}

		newsletterIDUint, err := strconv.ParseUint(newsletterID, 10, 64)
		if err != nil {
		    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Newsletter ID"})
		    return
		}

		newsletter.ID = uint(newsletterIDUint)

		err = newsletterService.UpdateNewsletter(&newsletter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Newsletter updated successfully"})
	}
}
