package routes

import (
	"go-clean/middlewares"
	"go-clean/modules/user"

	"github.com/gin-gonic/gin"
)

type NewUserRoutes struct{}

func (s *NewUserRoutes) Init(router *gin.Engine) {

	userRoutes := router.Group("/api/user")
	userRoutes.Use(middlewares.AuthMiddleware())
	{
		userRoutes.GET("/profile",
			middlewares.AuthMiddleware(),
			middlewares.Role(user.SUPERADMIN, user.ADMIN, user.USER),
			user.Controller.GetProfile)
		userRoutes.GET("/users",
			middlewares.AuthMiddleware(),
			middlewares.Role(user.SUPERADMIN),
			user.Controller.GetUsers)
		userRoutes.GET("/users/:id",
			middlewares.AuthMiddleware(),
			middlewares.Role(user.SUPERADMIN),
			middlewares.ValidationMiddleware(user.RequestQueryUser{}, nil),
			user.Controller.GetUser)
		userRoutes.POST("/users",
			middlewares.AuthMiddleware(),
			middlewares.Role(user.SUPERADMIN),
			middlewares.ValidationMiddleware(nil, &user.RequestCreateUser{}),
			user.Controller.CreateUser)
		userRoutes.PATCH("/users/:id",
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
