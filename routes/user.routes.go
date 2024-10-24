package routes

import (
	"github.com/gin-gonic/gin"
	"go-clean/middlewares"
	"go-clean/modules/user"
)

type NewUserRoutes struct{}

func (s *NewUserRoutes) Init(router *gin.Engine) {

	userRoutes := router.Group("/api/user")
	//userRoutes.Use(middlewares.AuthMiddleware())
	{
		userRoutes.GET("/profile",
			//middlewares.AuthMiddleware(),
			//middlewares.Role(auth.SUPERADMIN, auth.ADMIN, auth.USER),
			user.Controller.GetProfile)
		userRoutes.GET("/users",
			//middlewares.AuthMiddleware(),
			//middlewares.Role(auth.SUPERADMIN),
			user.Controller.GetUsers)
		userRoutes.GET("/users/:id",
			middlewares.AuthMiddleware(),
			middlewares.Role(user.SUPERADMIN),
			middlewares.ValidationMiddleware(user.RequestQueryUser{}, nil),
			user.Controller.GetUser)
		userRoutes.PUT("/users/:id",
			middlewares.AuthMiddleware(),
			middlewares.Role(user.SUPERADMIN),
			user.Controller.UpdateUSer)
		userRoutes.DELETE("/users/:id",
			middlewares.AuthMiddleware(),
			middlewares.Role(user.SUPERADMIN),
			user.Controller.DeleteUser)
	}
}

var UserRoute = &NewUserRoutes{}
