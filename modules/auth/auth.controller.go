package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewAuthController struct{}

func (s *NewAuthController) LoginEmail(c *gin.Context) {
	var req RequestLoginEmail
	_ = c.BindJSON(&req)

	userData, errCredential := AuthService.LoginWithUsernameAndEmail(req.Credential, req.Password)

	if errCredential != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  http.StatusUnauthorized,
			"error": errCredential.Error()})
		return
	}

	token, tokenErr := TokenService.GenerateToken(userData)

	if tokenErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":   http.StatusUnauthorized,
			"errors": "Could Not Generate Token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"user":  userData,
			"token": token,
		},
	})
}

func (s *NewAuthController) Login(c *gin.Context) {
	var req RequestLogin
	_ = c.BindJSON(&req)

	userData, errCredential := AuthService.LoginWithUsernameAndPassword(req.Username, req.Password)

	if errCredential != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  http.StatusUnauthorized,
			"error": errCredential.Error()})
		return
	}

	token, tokenErr := TokenService.GenerateToken(userData)

	if tokenErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":   http.StatusUnauthorized,
			"errors": "Could Not Generate Token",
		})
		return
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
			"code":   http.StatusUnauthorized,
			"errors": "Please authenticate",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Successfully logged out"})
}

func (s *NewAuthController) RefreshToken(c *gin.Context) {
	var req RequestRefreshToken
	_ = c.BindJSON(&req)
	generateRefreshAuth, err := AuthService.RefreshAuth(req.RefreshToken)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":   http.StatusUnauthorized,
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": generateRefreshAuth,
	})
}

var Controller = &NewAuthController{}
