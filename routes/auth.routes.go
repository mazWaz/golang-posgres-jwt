package routes

import (
	"go-clean/middlewares"
	"go-clean/modules/auth"

	"github.com/gin-gonic/gin"
)

type NewAuthRoutes struct{}

func (s *NewAuthRoutes) init(router *gin.Engine) {
	publicRoutes := router.Group("/api")
	{
		publicRoutes.POST("/login/admin",
			middlewares.ValidationMiddleware(auth.ValidateLoginEmail),
			auth.Controller.LoginEmail)
		publicRoutes.POST("/login",
			middlewares.ValidationMiddleware(auth.ValidateLogin),
			auth.Controller.Login)
		publicRoutes.POST("/refresh-token",
			middlewares.ValidationMiddleware(auth.ValidateRefreshToken),
			auth.Controller.RefreshToken)
		publicRoutes.POST(""+"/logout",
			middlewares.AuthMiddleware(),
			middlewares.ValidationMiddleware(auth.ValidateLogin),
			auth.Controller.Logout)
	}
}

var AuthRoutes = NewAuthRoutes{}
