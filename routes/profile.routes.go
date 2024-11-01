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
			middlewares.Role(middlewares.SUPERADMIN, middlewares.ADMIN, middlewares.USER),
			user.Controller.GetProfile)

		addressRoutes.GET("/",
			middlewares.Role(middlewares.SUPERADMIN, middlewares.ADMIN, middlewares.USER),
			middlewares.ValidationMiddleware(profile.ValidateQueryAddress),
			profile.Controller.GetProfile)

		addressRoutes.GET("/:id",
			middlewares.AuthMiddleware(),
			middlewares.Role(middlewares.SUPERADMIN, middlewares.ADMIN, middlewares.USER),
			profile.Controller.GetProfile)

		addressRoutes.POST("/",
			middlewares.Role(middlewares.SUPERADMIN, middlewares.ADMIN, middlewares.USER),
			middlewares.ValidationMiddleware(profile.ValidateCreateAddress),
			profile.Controller.CreateAddress)

		addressRoutes.PATCH("/:id",
			middlewares.Role(middlewares.SUPERADMIN, middlewares.ADMIN, middlewares.USER),
			middlewares.ValidationMiddleware(profile.ValidateUpdateAddress),
			profile.Controller.UpdateAddress)

		addressRoutes.DELETE("/:id",
			middlewares.Role(middlewares.SUPERADMIN, middlewares.ADMIN, middlewares.USER),
			middlewares.ValidationMiddleware(profile.ValidateDeleteAddress),
			profile.Controller.DeleteAddress)
	}
}

var AddressRoute = &NewAddressRoutes{}
