package routes

import (
	"go-clean/middlewares"
	"go-clean/modules/profile"
	"go-clean/modules/user"

	"github.com/gin-gonic/gin"
)

type NewAddressRoutes struct{}

func (s *NewAddressRoutes) Init(router *gin.Engine) {

	addressRoutes := router.Group("/api/profile")
	addressRoutes.Use(middlewares.AuthMiddleware())
	{
		addressRoutes.GET("/profile",
			middlewares.AuthMiddleware(),
			middlewares.Role(user.SUPERADMIN, user.ADMIN, user.USER),
			user.Controller.GetProfile)

		addressRoutes.GET("/",
			middlewares.Role(user.SUPERADMIN, user.ADMIN, user.USER),
			middlewares.ValidationMiddleware(&user.RequestQueryUser{}, nil),
			profile.Controller.GetProfile)

		addressRoutes.GET("/:id",
			middlewares.AuthMiddleware(),
			middlewares.Role(user.SUPERADMIN, user.ADMIN, user.USER),
			profile.Controller.GetProfile)

		addressRoutes.POST("/",
			middlewares.Role(user.SUPERADMIN, user.ADMIN, user.USER),
			middlewares.ValidationMiddleware(nil, &profile.RequestCreateAddress{}),
			profile.Controller.CreateAddress)

		addressRoutes.PATCH("/:id",
			middlewares.Role(user.SUPERADMIN, user.ADMIN, user.USER),
			middlewares.ValidationMiddleware(nil, &user.RequestUpdateUser{}),
			profile.Controller.UpdateAddress)

		addressRoutes.DELETE("/:id",
			middlewares.Role(user.SUPERADMIN, user.ADMIN, user.USER),
			profile.Controller.DeleteAddress)
	}
}

var AddressRoute = &NewAddressRoutes{}
