package user

import (
	"github.com/gin-gonic/gin"
	"go-clean/utils"
	"net/http"
	"strconv"
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
	id, err := strconv.Atoi(c.Param("id"))

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

func (s *NewController) UpdateUSer(c *gin.Context) {
	//id, err := strconv.Atoi(c.Param("id"))
}

func (s *NewController) DeleteUser(c *gin.Context) {}

var Controller = &NewController{}
