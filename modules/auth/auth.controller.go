package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type NewAuthController struct{}

func (s *NewAuthController) Login(c *gin.Context) {
	var req RequestLogin
	_ = c.BindJSON(&req)

	userData, errCredential := AuthService.LoginWithUsernameAndPassword(req.Username, req.Password)

	if errCredential != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errCredential})
	}

	token, tokenErr := TokenService.GenerateToken(userData)

	if tokenErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Could Not Generate Token",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"user":  userData,
			"token": token,
		},
	})
}

func (s *NewAuthController) Logout(c *gin.Context) {
	var req RequestRefreshToken
	_ = c.BindJSON(&req)
	logoutData := AuthService.LogoutWithRefreshToken(req.RefreshToken)

	if logoutData != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Please authenticate",
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func (s *NewAuthController) RefreshToken(c *gin.Context) {
	var req RequestRefreshToken
	_ = c.BindJSON(&req)

	generateRefreshAuth, err := AuthService.RefreshAuth(req.RefreshToken)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Please authenticate",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": generateRefreshAuth,
	})
}

var Controller = &NewAuthController{}
