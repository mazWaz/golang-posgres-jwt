package profile

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NewController struct{}

// NewAuthController is a constructor for AuthController

func (s *NewController) GetProfile(c *gin.Context) {
	id, _ := c.Get("userID")
	userId, _ := id.(float64)

	addressData, _ := Service.GetAddressByUserID(uint(userId))

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

func (s *NewController) CreateAddress(c *gin.Context) {
	id, _ := c.Get("userID")
	userId, _ := id.(float64)

	if id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code":  http.StatusNotFound,
			"error": "Failed to get id user",
		})
	}
	// Get form data
	var req RequestCreateAddress
	err := c.BindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"errors": "faild to get form data",
		})
		return
	}

	// Insert Data
	data, err := Service.CreateAddress(req, userId)

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
		"data": data,
	})
}

func (s *NewController) UpdateAddress(c *gin.Context) {
	// Get Id Address
	id, errId := strconv.Atoi(c.Param("id"))

	if errId != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":   http.StatusNotFound,
			"errors": "missing record id",
		})
		return
	}

	// Get form data
	var req RequestUpdateAddress
	err := c.BindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"errors": "cannot get form data",
		})
		return
	}

	addressId := uint(id)

	// Update data to db
	errUpdate := Service.UpdateAddress(addressId, req)

	if errUpdate != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"errors": "failed to update record",
		})
		return
	}

	// Get updated data in JSON
	updatedUser, errFetch := Service.GetAddressByID(addressId)
	if errFetch != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":   http.StatusInternalServerError,
			"errors": "failed to fetch updated record",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": updatedUser,
	})
}

func (s *NewController) DeleteAddress(c *gin.Context) {
	// Get record id
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"errors": "missing record id",
		})
	}
	addressId := uint(id)

	// Find record by id
	_, errFetch := Service.GetAddressByID(addressId)
	if errFetch != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
			"code":   http.StatusNonAuthoritativeInfo,
			"errors": "record doesn't exist",
		})
		return
	}

	// Delete record
	errDelete := Service.DeleteAddress(addressId)
	if errDelete != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"errors": "failed to delete record",
		})
		return
	}

	// Return
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "record has been deleted",
	})
}

var Controller = &NewController{}
