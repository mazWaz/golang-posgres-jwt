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

func (s *NewController) GetUsers(c *gin.Context) {

	var query RequestQueryUserByAdmin
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
	role, _ := c.Get("role")
	userRole, _ := role.(string)

	var userData RequestCreateUser
	_ = c.BindJSON(&userData)

	// Insert Data
	createdUser, errCreate := Service.CreateUser(userRole, &userData)

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
		"data": createdUser,
	})
}

func (s *NewController) UpdateUser(c *gin.Context) {
	// Get record id
	id, _ := strconv.Atoi(c.Param("id"))
	role, _ := c.Get("role")
	userRole, _ := role.(string)

	// Get form data
	var req RequestUpdateUser
	userId := uint(id)
	_ = c.BindJSON(&req)

	updatedUser, errUpdate := Service.UpdateUser(userRole, userId, &req)

	if errUpdate != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"errors": errUpdate.Error(),
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
	id, _ := strconv.Atoi(c.Param("id"))
	role, _ := c.Get("role")
	userId := uint(id)
	userRole, _ := role.(string)

	// TODO: Delete Adress
	// Delete record
	errDelete := Service.DeleteUser(userRole, userId)
	if errDelete != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"errors": errDelete.Error(),
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
