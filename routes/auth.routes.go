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
		publicRoutes.POST("/login",
			// middlewares.ValidationMiddleware(nil, auth.RequestLogin{}),
			auth.Controller.Login)
		publicRoutes.POST("/refresh-token",
			middlewares.ValidationMiddleware(nil, auth.RequestRefreshToken{}),
			auth.Controller.RefreshToken)
		publicRoutes.POST(""+"/logout",
			middlewares.AuthMiddleware(),
			middlewares.ValidationMiddleware(nil, auth.RequestRefreshToken{}),
			auth.Controller.Logout)
	}
}

var AuthRoutes = NewAuthRoutes{}
