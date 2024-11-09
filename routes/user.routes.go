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
		userRoutes.GET("/",
			middlewares.Role(middlewares.SUPERADMIN, middlewares.ADMIN),
			middlewares.ValidationMiddleware(user.ValidateQueryUser),
			user.Controller.GetUsers)

		userRoutes.GET("/:id",
			middlewares.Role(middlewares.SUPERADMIN, middlewares.ADMIN),
			user.Controller.GetUser)

		userRoutes.PATCH("/:id",
			middlewares.Role(middlewares.SUPERADMIN, middlewares.ADMIN),
			middlewares.ValidationMiddleware(user.ValidateUpdateUser),
			user.Controller.UpdateUser)

		userRoutes.DELETE("/:id",
			middlewares.Role(middlewares.SUPERADMIN, middlewares.ADMIN),
			middlewares.ValidationMiddleware(user.ValidateDeleteUser),
			user.Controller.DeleteUser)

		userRoutes.POST("/",
			middlewares.Role(middlewares.SUPERADMIN, middlewares.ADMIN),
			middlewares.ValidationMiddleware(user.ValidateCreateUser),
			user.Controller.CreateUser)
	}
}

var UserRoute = &NewUserRoutes{}
