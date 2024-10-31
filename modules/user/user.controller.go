package user

import (
	"errors"
	"go-clean/utils"
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type NewController struct{}

// NewAuthController is a constructor for AuthController

func (s *NewController) GetProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	userData, err := Service.GetUserByID(uint(id))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"code":   http.StatusNotFound,
				"errors": "User Not Found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":   http.StatusInternalServerError,
			"errors": "Internal server error",
		})
		return
	}

	if userData == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":   http.StatusNotFound,
			"errors": "User Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": userData,
	})
}

func (s *NewController) GetUsers(c *gin.Context) {

	var query RequestQueryUser
	_ = c.Bind(&query)

	filters := map[string]interface{}{
		"username LIKE ?": "%" + query.Username + "%",
		"role = ?":        query.Role,
	}

	var users []ModelUser

	pagination, err := utils.Paginate(
		int(query.Page),
		int(query.Limit),
		filters,
		&ModelUser{},
		&users,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":   http.StatusInternalServerError,
			"errors": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"users":      pagination.Data,
			"pagination": pagination.Pagination,
		},
	})

}

func (s *NewController) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	userData, err := Service.GetUserByID(uint(id))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"code":   http.StatusNotFound,
				"errors": "User Not Found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":   http.StatusInternalServerError,
			"errors": "Internal server error",
		})
		return
	}

	if userData == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":   http.StatusNotFound,
			"errors": "User Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": userData,
	})
}

func (s *NewController) CreateUser(c *gin.Context) {
	// Get form data
	var req RequestCreateUser

	_ = c.BindJSON(&req)

	hashedPassword, err := utils.HashPassword([]byte(req.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"errors": "failed to generate password",
		})
		return
	}

	// Assign form data JSON to struct
	var userData ModelUser
	userData.Username = req.Username
	userData.Password = string(hashedPassword)
	userData.Role = req.Role

	// Insert Data
	errCreate := Service.CreateUser(&userData)

	if errCreate != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"errors": errCreate.Error(),
		})
		return
	}

	// Return data
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": userData,
	})
}

func (s *NewController) UpdateUSer(c *gin.Context) {
	// Get record id
	id, errId := strconv.Atoi(c.Param("id"))

	if errId != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":   http.StatusNotFound,
			"errors": "missing record id",
		})
		return
	}

	// Get form data
	var req RequestUpdateUser
	err := c.BindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"errors": "cannot get form data",
		})
		return
	}
	userId := uint(id)

	// Update data to db
	errUpdate := Service.UpdateUser(userId, req)

	if errUpdate != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"errors": "failed to update record",
		})
		return
	}

	// Get updated data in JSON
	updatedUser, errFetch := Service.GetUserByID(userId)
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

func (s *NewController) DeleteUser(c *gin.Context) {
	// Get record id
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"errors": "missing record id",
		})
	}
	userId := uint(id)

	// Find record by id
	_, errFetch := Service.GetUserByID(userId)
	if errFetch != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
			"code":   http.StatusNonAuthoritativeInfo,
			"errors": "record doesn't exist",
		})
		return
	}

	// Delete record
	errDelete := Service.DeleteUser(userId)
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
