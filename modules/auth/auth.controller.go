package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var req ValidateLogin
	_ = c.BindJSON(&req)

	userData, errCredential := LoginWithUsernameAndPassword(req.Username, req.Password)

	if errCredential != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errCredential})
	}

	token, tokenErr := generateToken(userData)

	if tokenErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Could Not Generate Token",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"user":  userData,
		"token": token,
	})
}

func RefreshToken(c *gin.Context) {
	var req ValidateRefreshToken
	_ = c.BindJSON(&req)

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func Logout(c *gin.Context) {
	userInterface, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	user := userInterface.(models.User)
	user.TokenVersion += 1
	if err := models.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not logout user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
