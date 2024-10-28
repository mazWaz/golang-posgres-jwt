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
			middlewares.Role(user.SUPERADMIN, user.ADMIN, user.USER),
			user.Controller.GetProfile)

		userRoutes.GET("/",
			middlewares.Role(user.SUPERADMIN),
			middlewares.ValidationMiddleware(&user.RequestQueryUser{}, nil),
			user.Controller.GetUsers)

		userRoutes.GET("/:id",
			middlewares.Role(user.SUPERADMIN),
			user.Controller.GetUser)

		userRoutes.POST("/",
			middlewares.Role(user.SUPERADMIN),
			middlewares.ValidationMiddleware(nil, &user.RequestCreateUser{}),
			user.Controller.CreateUser)

		userRoutes.PATCH("/:id",
			middlewares.Role(user.SUPERADMIN),
			middlewares.ValidationMiddleware(nil, &user.RequestUpdateUser{}),
			user.Controller.UpdateUSer)

		userRoutes.DELETE("/:id",
			middlewares.Role(user.SUPERADMIN),
			user.Controller.DeleteUser)
	}
}

var UserRoute = &NewUserRoutes{}
