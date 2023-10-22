package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go_newsletter_api/internal/auth"
	"go_newsletter_api/internal/user/service"
)

func LoginHandler(c *gin.Context, userService *service.UserService) {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	user, err := userService.GetUserByEmail(loginRequest.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(401, gin.H{"error": "User not found"})
		} else {
			c.JSON(500, gin.H{"error": "Failed to check user"})
		}
		return
	}

	if err := userService.VerifyPassword(user.Password, loginRequest.Password); err != nil {
		c.JSON(401, gin.H{"error": "Invalid password"})
		return
	}

	token, err := auth.GenerateToken(user.ID, user.Role)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetCookie("auth_token", token, 3600, "/", "localhost", false, true)

	c.JSON(200, gin.H{"message": "Login successful", "token": token})
}
