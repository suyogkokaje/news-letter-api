package handlers

import (
    "github.com/gin-gonic/gin"
    "go_newsletter_api/internal/user/model"
    "go_newsletter_api/internal/user/service"
    "github.com/jinzhu/gorm"
)

func SignUpHandler(c *gin.Context, userService *service.UserService) {
    var user model.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": "Invalid request data"})
        return
    }

    existingUser, err := userService.GetUserByEmail(user.Email)
    if err != nil && err != gorm.ErrRecordNotFound {
        c.JSON(500, gin.H{"error": "Failed to check existing user"})
        return
    }
    
    if existingUser != nil {
        c.JSON(400, gin.H{"error": "A user with this email already exists"})
        return
    }

    if user.Role == "" {
        user.Role = "user"
    }

    newUser, err := userService.CreateUser(&user)
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(200, gin.H{"message": "User created successfully", "user": newUser})
}
