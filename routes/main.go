package routes

import (
	"github.com/gin-gonic/gin"
	"go-clean/middlewares"
	"go-clean/modules/auth"
	"go-clean/modules/user"
)

func SetupRoutes(router *gin.Engine) {
	publicRoutes := router.Group("/api")
	{
		publicRoutes.POST("/login",
			middlewares.ValidationMiddleware(nil, auth.RequestLogin{}),
			auth.Login)
		publicRoutes.POST("/refresh-token",
			middlewares.ValidationMiddleware(nil, auth.RequestRefreshToken{}),
			auth.RefreshToken)
		publicRoutes.POST(""+"/logout",
			middlewares.AuthMiddleware(),
			middlewares.ValidationMiddleware(nil, auth.RequestRefreshToken{}),
			auth.Logout)
	}

	userRoutes := router.Group("/api/user")
	userRoutes.Use(middlewares.AuthMiddleware())
	{
		userRoutes.GET("/profile",
			middlewares.AuthMiddleware(),
			middlewares.Role(auth.SUPERADMIN, auth.ADMIN, auth.USER),
			controllers.GetProfile)
		// Other user-specific routes
	}

	adminRoutes := router.Group("/api/admin")
	adminRoutes.Use()
	{
		adminRoutes.GET("/users",
			middlewares.AuthMiddleware(),
			middlewares.Role(auth.SUPERADMIN),
			user.GetUsers)
		adminRoutes.GET("/users/:id",
			middlewares.AuthMiddleware(),
			middlewares.Role(auth.SUPERADMIN),
			user.GetUser)
		adminRoutes.PUT("/users/:id",
			middlewares.AuthMiddleware(),
			middlewares.Role(auth.SUPERADMIN),
			controllers.UpdateUser)
		adminRoutes.DELETE("/users/:id",
			middlewares.AuthMiddleware(),
			middlewares.Role(auth.SUPERADMIN),
			controllers.DeleteUser)
	}
}
