package profile

import (
	"go-clean/modules/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewController struct{}

// NewAuthController is a constructor for AuthController
func (s *NewController) GetProfile(c *gin.Context) {
	id, _ := c.Get("userID")
	userId, _ := id.(float64)

	addressData, _ := user.Service.GetFullUserByUserID(uint(userId))

	if addressData == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":   http.StatusNotFound,
			"errors": "Address Not Found",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": addressData,
	})
}

func (s *NewController) UpdateAddress(c *gin.Context) {
	id, _ := c.Get("userID")
	userId, _ := id.(float64)
	role, _ := c.Get("role")
	userRole, _ := role.(string)

	// Get form data
	var req user.RequestUpdateAddress
	_ = c.BindJSON(&req)

	if userId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "not logged in",
		})
	}

	// Update Data
	ResponseData, err := user.Service.UpdateAddress(uint(userId), userRole, &req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"errors": err.Error(),
		})
		return
	}

	// Return data
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"data": ResponseData,
		},
	})
}

var Controller = &NewController{}
