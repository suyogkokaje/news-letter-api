package handlers

import (
    "go_newsletter_api/internal/auth"
    "go_newsletter_api/internal/news_letter/service"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

func DeleteNewsletterHandler(newsletterService *service.NewsletterService) gin.HandlerFunc {
    return func(c *gin.Context) {
        newsletterIDStr := c.Param("id")
        newsletterID, err := strconv.ParseUint(newsletterIDStr, 10, 32)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid newsletter ID"})
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

        err = newsletterService.DeleteNewsletter(uint(newsletterID))
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Newsletter not found"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Newsletter deleted"})
    }
}
