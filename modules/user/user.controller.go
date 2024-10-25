package user

import (
	"go-clean/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type NewController struct{}

// NewAuthController is a constructor for AuthController

func (s *NewController) GetProfile(c *gin.Context) {

}

func (s *NewController) GetUsers(c *gin.Context) {

	var query RequestQueryUser
	_ = c.BindQuery(&query)

	filters := map[string]interface{}{
		"username LIKE": "%" + query.Username + "%",
		"role":          query.Role,
	}

	var users []ModelUser

	pagination, err := utils.Paginate(
		1,
		100,
		filters,
		&ModelUser{},
		&users,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Internal server error",
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Internal server error",
		})
		return
	}

	if userData == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "User Not Found",
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

	//Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to generate password",
		})
		return
	}

	// Assign form data JSON to struct
	var userData ModelUser
	userData.Username = req.Username
	userData.Password = string(hash)
	userData.Role = req.Role

	// Insert Data
	errCreate := Service.CreateUser(&userData)

	if errCreate != nil {
		c.JSON(400, gin.H{
			"error": "failed to store data in database",
		})
		return
	}

	// Return data
	c.JSON(200, gin.H{
		"data": userData,
	})
}

func (s *NewController) UpdateUSer(c *gin.Context) {
	// Get record id
	id, errId := strconv.Atoi(c.Param("id"))

	if errId != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "missing record id",
		})
	}

	// Get form data
	var req RequestUpdateUser
	err := c.BindJSON(&req)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot get form data",
		})
		return
	}
	userId := uint(id)

	// Update data to db
	errUpdate := Service.UpdateUser(userId, req)

	if errUpdate != nil {
		c.JSON(400, gin.H{
			"error": "failed to update record",
		})
		return
	}

	// Get updated data in JSON
	updatedUser, errFetch := Service.GetUserByID(userId)
	if errFetch != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch updated record",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "record has been updated",
		"data":    updatedUser,
	})
}

func (s *NewController) DeleteUser(c *gin.Context) {
	// Get record id
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(404, gin.H{
			"error": "missing record id",
		})
	}
	userId := uint(id)

	// Find record by id
	_, errFetch := Service.GetUserByID(userId)
	if errFetch != nil {
		c.JSON(500, gin.H{
			"error": "record doesn't exist",
		})
		return
	}

	// Delete record
	errDelete := Service.DeleteUser(userId)
	if errDelete != nil {
		c.JSON(400, gin.H{
			"error": "failed to delete record",
		})
		return
	}

	// Return
	c.JSON(200, gin.H{
		"message": "record has been deleted",
	})
}

var Controller = &NewController{}
