package routes

import (
    "github.com/gin-gonic/gin"
    "go_newsletter_api/internal/user/handlers"
    "go_newsletter_api/internal/user/service"
)

func SetupRouter(r *gin.Engine, userService *service.UserService) {
    userRoutes := r.Group("/user")
    {
        userRoutes.POST("/signup", func(c *gin.Context) {
            handlers.SignUpHandler(c, userService)
        })

        userRoutes.POST("/login", func(c *gin.Context) {
            handlers.LoginHandler(c, userService)
        })
    }
}
