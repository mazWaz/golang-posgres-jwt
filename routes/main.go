package routes

import (
	"github.com/gin-gonic/gin"
	"go-clean/middlewares"
)

func SetupRoutes(router *gin.Engine) {
	publicRoutes := router.Group("/api")
	{
		publicRoutes.POST("/login", controllers.Login)
		publicRoutes.POST("/refresh-token", controllers.RefreshToken)
		// Add register route if needed
	}

	userRoutes := router.Group("/api/user")
	userRoutes.Use(middlewares.AuthMiddleware())
	{
		userRoutes.GET("/profile", controllers.GetProfile)
		userRoutes.POST("/logout", controllers.Logout)
		// Other user-specific routes
	}

	adminRoutes := router.Group("/api/admin")
	adminRoutes.Use(middlewares.AuthMiddleware(), middlewares.AdminOnly())
	{
		adminRoutes.GET("/users", controllers.GetAllUsers)
		adminRoutes.GET("/users/:id", controllers.GetUserByID)
		adminRoutes.PUT("/users/:id", controllers.UpdateUser)
		adminRoutes.DELETE("/users/:id", controllers.DeleteUser)
	}
}
