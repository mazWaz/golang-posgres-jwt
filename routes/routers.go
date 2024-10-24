package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	AuthRoutes.init(r)
	UserRoute.Init(r)
}
