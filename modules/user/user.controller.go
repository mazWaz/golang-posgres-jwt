package user

import (
	"fmt"
	"go-clean/db"
	"go-clean/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	fmt.Println(c.Request.Body)

	_ = c.BindJSON(&req)

	// Assign form data JSON to struct
	var userData ModelUser
	userData.Username = req.Username
	userData.Password = req.Password
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
	var userData ModelUser
	var req RequestUpdateUser

	err := c.BindJSON(&req)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot get form data",
		})
		return
	}

	// Find record by id
	errFind := db.Data.Table("users").Where("id = ?", id).Find(&userData).Error
	if errFind != nil {
		c.JSON(500, gin.H{
			"error": "internal server error",
		})
		return
	}

	if userData.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
	}

	userId := uint(id)

	// Update data to db
	errUpdate := Service.UpdateUser(userId, req, &userData)

	if errUpdate != nil {
		c.JSON(400, gin.H{
			"error": "failed to update record",
		})
		return
	}

	// return
	c.JSON(200, gin.H{
		"message": "record has been updated",
		"data":    userData,
	})
}

func (s *NewController) DeleteUser(c *gin.Context) {
	var userData ModelUser

	// Get record id
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(404, gin.H{
			"error": "missing record id",
		})
	}

	// Find record by id
	errFind := db.Data.Table("users").Where("id = ?", id).Find(&userData).Error
	if errFind != nil {
		c.JSON(500, gin.H{
			"error": "internal server error",
		})
		return
	}

	if userData.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
	}

	// Delete record
	userId := uint(id)
	errDelete := Service.DeleteUser(userId)

	if errDelete != nil {
		c.JSON(400, gin.H{
			"error": "failed to delete record",
		})
		return
	}

	// Return
	c.JSON(200, gin.H{

		"error": "record has been deleted",
	})
}

var Controller = &NewController{}
