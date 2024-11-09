package routes

import (
	"go-clean/middlewares"
	"go-clean/modules/profile"

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
			profile.Controller.GetProfile)

		addressRoutes.GET("/",
			middlewares.Role(middlewares.SUPERADMIN, middlewares.ADMIN, middlewares.USER),
			middlewares.ValidationMiddleware(profile.ValidateQueryAddress),
			profile.Controller.GetProfile)

		addressRoutes.GET("/:id",
			middlewares.AuthMiddleware(),
			middlewares.Role(middlewares.SUPERADMIN, middlewares.ADMIN, middlewares.USER),
			profile.Controller.GetProfile)

		addressRoutes.PATCH("/:id",
			middlewares.Role(middlewares.SUPERADMIN, middlewares.ADMIN, middlewares.USER),
			middlewares.ValidationMiddleware(profile.ValidateUpdateAddress),
			profile.Controller.UpdateAddress)
	}
}

var AddressRoute = &NewAddressRoutes{}
